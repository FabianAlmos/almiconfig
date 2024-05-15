package almiconfig

import (
	"github.com/FabianAlmos/almiconfig/consts"
	almierrors "github.com/FabianAlmos/almiconfig/errors"
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

func setFieldValue(envVar any, cfg reflect.Value, val *configValue, cc *configConstraint) error {
	envVarValue := reflect.ValueOf(envVar)
	field := cfg.FieldByName(val.Field.Name)
	structTagType := string(regexp.MustCompile(slice).ReplaceAll([]byte(field.Type().String()), []byte(consts.EMPTY)))

	if !(cc.SliceType && cc.Type == structTagType) &&
		field.Type().String() != cc.Type &&
		!(field.Type().String() == _string && cc.Type == consts.EMPTY) {
		return almierrors.FieldStructTagTypeMismatchErr.Build(
			val.Field.Name,
			field.Type().String(),
			cfg.Type().String(),
			cc.Type,
			cfg.Type().String(),
		)
	}

	field.Set(envVarValue)

	return nil
}

func ValidateConfig[T any](config T) (*T, error) {
	cfg := reflect.ValueOf(&config).Elem()
	for i := 0; i < cfg.NumField(); i++ {
		val := getConfigValueByIndex(cfg, i)

		cfgConstraint := newConfigConstraint(val)
		if err := cfgConstraint.parseConstraints(val.Constraints); err != nil {
			return nil, err
		}

		envVar, err := cfgConstraint.findType()
		if err != nil {
			return nil, almierrors.FailedToConvertTypeErr.Build(cfgConstraint.EnvName, cfgConstraint.Type)
		}

		if err := setFieldValue(envVar, cfg, val, cfgConstraint); err != nil {
			return nil, err
		}

		if err := cfgConstraint.checkConstraints(val); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
