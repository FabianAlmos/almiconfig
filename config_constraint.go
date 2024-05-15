package almi

import (
	"almi/consts"
	almierrors "almi/errors"
	"regexp"
)

type configConstraint struct {
	FieldName string

	Required bool
	EnvName  string
	Type     string

	SliceType bool
	Separator string
}

func newConfigConstraint(val *configValue) *configConstraint {
	return &configConstraint{
		FieldName: val.Field.Name,
	}
}

func (cc *configConstraint) parseConstraints(constraints []string) error {
	for _, c := range constraints {
		switch {
		case regexp.MustCompile(required).MatchString(c):
			cc.Required = true
			continue
		case regexp.MustCompile(env).MatchString(c):
			cc.EnvName = string(regexp.MustCompile(envEq).ReplaceAll([]byte(c), []byte("")))
			continue
		case regexp.MustCompile(_type).MatchString(c):
			if regexp.MustCompile(typeSlice).MatchString(c) {
				sliceType := regexp.MustCompile(typeEq).ReplaceAll([]byte(c), []byte(consts.EMPTY))
				sep := regexp.MustCompile(sliceSep).Find(sliceType)
				if len(sep) != 3 {
					return almierrors.SepUndefErr.Build(cc.FieldName)
				}
				cc.Type = string(regexp.MustCompile(sliceSep).ReplaceAll(sliceType, []byte(consts.EMPTY)))
				cc.SliceType = true
				cc.Separator = string(sep[1])
				continue
			}

			cc.Type = string(regexp.MustCompile(typeEq).ReplaceAll([]byte(c), []byte(consts.EMPTY)))
			continue
		default:
			return almierrors.ConstraintUnknownErr.Build(c, cc.FieldName)
		}
	}

	return nil
}

func (cc *configConstraint) findType() (any, error) {
	var (
		envVar any
		err    error
	)

	switch cc.Type {
	case consts.EMPTY, _string:
		envVar, err = str[string](*cc)
	case _bool:
		envVar, err = atob[bool](*cc)
	case _int:
		envVar, err = atoi[int](*cc)
	case _int8:
		envVar, err = atoi[int8](*cc)
	case _int16:
		envVar, err = atoi[int16](*cc)
	case _int32:
		envVar, err = atoi[int32](*cc)
	case _int64:
		envVar, err = atoi[int64](*cc)
	case _uint8:
		envVar, err = atoi[uint8](*cc)
	case _uint16:
		envVar, err = atoi[uint16](*cc)
	case _uint32:
		envVar, err = atoi[uint32](*cc)
	case _uint64:
		envVar, err = atoi[uint64](*cc)
	case _uintptr:
		envVar, err = atoi[uintptr](*cc)
	case _float32:
		envVar, err = atoi[float32](*cc)
	case _float64:
		envVar, err = atoi[float64](*cc)
	case _byte:
		envVar, err = atoRB[byte](*cc)
	case _rune:
		envVar, err = atoRB[rune](*cc)
	default:
		return nil, almierrors.UnrecognizedTypeErr.Build(cc.Type)
	}

	return envVar, err
}

func (cc *configConstraint) checkConstraints(val *configValue) error {
	if cc.EnvName == consts.EMPTY {
		return almierrors.EnvConstraintUndefErr.Build(val.Field.Name)
	}

	if cc.Required && val.Value.String() == consts.EMPTY {
		return almierrors.FieldRequiredErr.Build(val.Field.Name)
	}

	return nil
}
