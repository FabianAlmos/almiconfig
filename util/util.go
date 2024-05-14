package util

import (
	"almi/consts"
	almitypes "almi/types"
	"errors"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
)

type IntConstraint interface {
	constraints.Signed | constraints.Unsigned
}

type Number interface {
	IntConstraint | constraints.Float
}

func AlmiAtoi[T Number](cc almitypes.ConfigConstraint) (any, error) {
	envVal := os.Getenv(cc.EnvName)
	if !cc.Required && envVal == consts.EMPTY {
		return T(0), nil
	}

	if cc.SliceType && cc.Separator != consts.EMPTY {
		var ns []T
		vals := strings.Split(envVal, cc.Separator)
		for _, val := range vals {
			n, err := strconv.Atoi(val)
			if err != nil {
				return T(0), err
			}

			ns = append(ns, T(n))
		}

		return ns, nil
	} else if cc.SliceType && cc.Separator == consts.EMPTY {
		return T(0), errors.New("AlmiConfig: separator must be specified for AlmiAtoi func when 'val' is of type []T")
	}

	n, err := strconv.Atoi(envVal)
	if err != nil {
		return T(0), err
	}

	return T(n), nil
}

// Generic for readability and proper zero value return
func AlmiStr[T ~string](cc almitypes.ConfigConstraint) (val any, err error) {
	envVal := os.Getenv(cc.EnvName)
	if !cc.Required && envVal == consts.EMPTY {
		return T(consts.EMPTY), nil
	}

	if cc.SliceType && cc.Separator != consts.EMPTY {
		var strs []T
		vals := strings.Split(envVal, cc.Separator)
		for _, val := range vals {
			strs = append(strs, T(val))
		}

		return strs, nil
	} else if cc.SliceType && cc.Separator == consts.EMPTY {
		return T(consts.EMPTY), errors.New("AlmiConfig: separator must be specified for AlmiStr func when 'val' is of type []T")
	}

	return T(envVal), nil
}

// Generic for readability and proper zero value return
func AlmiAtob[T ~bool](cc almitypes.ConfigConstraint) (val any, err error) {
	envVal := os.Getenv(cc.EnvName)
	if !cc.Required && envVal == consts.EMPTY {
		return T(false), nil
	}

	if cc.SliceType && cc.Separator != consts.EMPTY {
		var bs []T
		vals := strings.Split(envVal, cc.Separator)
		for _, val := range vals {
			b, err := strconv.ParseBool(val)
			if err != nil {
				return b, err
			}

			bs = append(bs, T(b))
		}

		return bs, nil
	} else if cc.SliceType && cc.Separator == consts.EMPTY {
		return T(false), errors.New("AlmiConfig: separator must be specified for AlmiAtob func when 'val' is of type []T")
	}

	b, err := strconv.ParseBool(envVal)
	if err != nil {
		return T(false), err
	}

	return T(b), nil
}

func AlmiAtoRB[T ~rune | ~byte](cc almitypes.ConfigConstraint) (val any, err error) {
	envVal := os.Getenv(cc.EnvName)
	if !cc.Required && envVal == consts.EMPTY {
		return T(0), nil
	}

	if cc.SliceType && cc.Separator != consts.EMPTY {
		var rbs []T
		vals := strings.Split(envVal, cc.Separator)
		for _, val := range vals {
			n, err := strconv.Atoi(val)
			if err != nil {
				return T(0), errors.New("AlmiConfig: failed to convert string to int to convert to rune/byte")
			}

			rbs = append(rbs, T(n))
		}

		return rbs, nil
	} else if cc.SliceType && cc.Separator == consts.EMPTY {
		return T(0), errors.New("AlmiConfig: separator must be specified for AlmiAtoRB func when 'val' is of type []T")
	}

	n, err := strconv.Atoi(envVal)
	if err != nil {
		return T(0), errors.New("AlmiConfig: failed to convert string to int to convert to rune/byte")
	}

	return T(n), nil
}
