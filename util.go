package almiconfig

import (
	"github.com/FabianAlmos/almiconfig/consts"
	almierrors "github.com/FabianAlmos/almiconfig/errors"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
)

type intConstraint interface {
	constraints.Signed | constraints.Unsigned
}

type number interface {
	intConstraint | constraints.Float
}

func atoi[T number](cc configConstraint) (any, error) {
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
		return T(0), almierrors.SepAtoiErr.Build()
	}

	n, err := strconv.Atoi(envVal)
	if err != nil {
		return T(0), err
	}

	return T(n), nil
}

// Generic for readability and proper zero value return
func str[T ~string](cc configConstraint) (val any, err error) {
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
		return T(consts.EMPTY), almierrors.SepStrErr.Build()
	}

	return T(envVal), nil
}

// Generic for readability and proper zero value return
func atob[T ~bool](cc configConstraint) (val any, err error) {
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
		return T(false), almierrors.SepAtobErr.Build()
	}

	b, err := strconv.ParseBool(envVal)
	if err != nil {
		return T(false), err
	}

	return T(b), nil
}

func atoRB[T ~rune | ~byte](cc configConstraint) (val any, err error) {
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
				return T(0), almierrors.AtoRBConversionFailed.Build()
			}

			rbs = append(rbs, T(n))
		}

		return rbs, nil
	} else if cc.SliceType && cc.Separator == consts.EMPTY {
		return T(0), almierrors.SepAtoRBErr.Build()
	}

	n, err := strconv.Atoi(envVal)
	if err != nil {
		return T(0), almierrors.AtoRBConversionFailed.Build()
	}

	return T(n), nil
}
