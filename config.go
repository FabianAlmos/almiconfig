package almi

import (
	"almi/consts"
	"almi/lexer"
	almitypes "almi/types"
	"almi/util"
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

const (
	almi = "almi"

	required  = "^(required)$"
	envEq     = "(env=)"
	env       = "^(env=.+)$"
	typeEq    = "(type=)"
	_type     = "^(type=.+)$"
	sliceSep  = "\\[.{1}\\]"
	slice     = "\\[\\]"
	typeSlice = "^(type=\\[.{1}\\].+)$"

	_bool    = "bool"
	_string  = "string"
	_int     = "int"
	_int8    = "int8"
	_int16   = "int16"
	_int32   = "int32"
	_int64   = "int64"
	_uint8   = "uint8"
	_uint16  = "uint16"
	_uint32  = "uint32"
	_uint64  = "uint64"
	_uintptr = "uintptr"
	_float32 = "float32"
	_float64 = "float64"
	_rune    = "rune"
	_byte    = "byte"
)

func getConfigValueByIndex(cfg reflect.Value, i int) *almitypes.ConfigValue {
	field := cfg.Type().Field(i)
	tag := field.Tag.Get(almi)

	lxr := lexer.NewLexer(tag)

	return &almitypes.ConfigValue{
		Field:       field,
		Tag:         tag,
		Constraints: lxr.Tokenize(),
		Value:       cfg.FieldByIndex([]int{i}),
	}
}

func parseConstraints(constraint *almitypes.ConfigConstraint, constraints []string) error {
	for _, c := range constraints {
		switch {
		case regexp.MustCompile(required).MatchString(c):
			constraint.Required = true
			continue
		case regexp.MustCompile(env).MatchString(c):
			constraint.EnvName = string(regexp.MustCompile(envEq).ReplaceAll([]byte(c), []byte("")))
			continue
		case regexp.MustCompile(_type).MatchString(c):
			if regexp.MustCompile(typeSlice).MatchString(c) {
				sliceType := regexp.MustCompile(typeEq).ReplaceAll([]byte(c), []byte(consts.EMPTY))
				sep := regexp.MustCompile(sliceSep).Find(sliceType)
				if len(sep) != 3 {
					return errors.New(
						fmt.Sprintf(
							"AlmiConfig: field: '%s': slice types must specify a separator character in their brackets",
							constraint.FieldName,
						),
					)
				}
				constraint.Type = string(regexp.MustCompile(sliceSep).ReplaceAll(sliceType, []byte(consts.EMPTY)))
				constraint.SliceType = true
				constraint.Separator = string(sep[1])
				continue
			}

			constraint.Type = string(regexp.MustCompile(typeEq).ReplaceAll([]byte(c), []byte(consts.EMPTY)))
			continue
		default:
			return errors.New(
				fmt.Sprintf(
					"AlmiConfig: Constarint: '%s' at Field: '%s', is unknown to almi config",
					c,
					constraint.FieldName,
				),
			)
		}
	}

	return nil
}

func findType(cc almitypes.ConfigConstraint) (any, error) {
	var (
		envVar any
		err    error
	)

	switch cc.Type {
	case consts.EMPTY, _string:
		envVar, err = util.AlmiStr[string](cc)
	case _bool:
		envVar, err = util.AlmiAtob[bool](cc)
	case _int:
		envVar, err = util.AlmiAtoi[int](cc)
	case _int8:
		envVar, err = util.AlmiAtoi[int8](cc)
	case _int16:
		envVar, err = util.AlmiAtoi[int16](cc)
	case _int32:
		envVar, err = util.AlmiAtoi[int32](cc)
	case _int64:
		envVar, err = util.AlmiAtoi[int64](cc)
	case _uint8:
		envVar, err = util.AlmiAtoi[uint8](cc)
	case _uint16:
		envVar, err = util.AlmiAtoi[uint16](cc)
	case _uint32:
		envVar, err = util.AlmiAtoi[uint32](cc)
	case _uint64:
		envVar, err = util.AlmiAtoi[uint64](cc)
	case _uintptr:
		envVar, err = util.AlmiAtoi[uintptr](cc)
	case _float32:
		envVar, err = util.AlmiAtoi[float32](cc)
	case _float64:
		envVar, err = util.AlmiAtoi[float64](cc)
	case _byte:
		envVar, err = util.AlmiAtoRB[byte](cc)
	case _rune:
		envVar, err = util.AlmiAtoRB[rune](cc)
	default:
		return nil, errors.New(
			fmt.Sprintf(
				"AlmiConfig: unrecognized type: '%s'",
				cc.Type,
			),
		)
	}

	return envVar, err
}

func setFieldValue(envVar any, cfg reflect.Value, val *almitypes.ConfigValue, cc almitypes.ConfigConstraint) error {
	envVarValue := reflect.ValueOf(envVar)
	field := cfg.FieldByName(val.Field.Name)
	structTagType := string(regexp.MustCompile(slice).ReplaceAll([]byte(field.Type().String()), []byte(consts.EMPTY)))

	if !(cc.SliceType && cc.Type == structTagType) &&
		field.Type().String() != cc.Type &&
		!(field.Type().String() == _string && cc.Type == consts.EMPTY) {
		return errors.New(
			fmt.Sprintf(
				"AlmiConfig: field: '%s' type: '%s' in config struct does not match the constraint type: '%s' in config struct tag",
				val.Field.Name,
				field.Type().String(),
				cc.Type,
			),
		)
	}

	field.Set(envVarValue)

	return nil
}

func checkConstraints(cc almitypes.ConfigConstraint, val *almitypes.ConfigValue) error {
	if cc.EnvName == consts.EMPTY {
		return errors.New(
			fmt.Sprintf(
				"AlmiConfig: 'env=' constraint must be defined for all fields of the config, constraint not found for field: '%s'",
				val.Field.Name,
			),
		)
	}

	if cc.Required && val.Value.String() == consts.EMPTY {
		return errors.New(
			fmt.Sprintf(
				"AlmiConfig: Field: '%s', is required",
				val.Field.Name,
			),
		)
	}

	return nil
}

func ValidateConfig[T any](config T) (*T, error) {
	cfg := reflect.ValueOf(&config).Elem()
	for i := 0; i < cfg.NumField(); i++ {
		val := getConfigValueByIndex(cfg, i)

		cfgConstraint := almitypes.ConfigConstraint{FieldName: val.Field.Name}
		if err := parseConstraints(&cfgConstraint, val.Constraints); err != nil {
			return nil, err
		}

		envVar, err := findType(cfgConstraint)
		if err != nil {
			return nil, errors.New(
				fmt.Sprintf(
					"AlmiConfig: failed to convert type of '%s' to %s from string",
					cfgConstraint.EnvName,
					cfgConstraint.Type,
				),
			)
		}

		if err := setFieldValue(envVar, cfg, val, cfgConstraint); err != nil {
			return nil, err
		}

		if err := checkConstraints(cfgConstraint, val); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
