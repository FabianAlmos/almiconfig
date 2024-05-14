package almierrors

import (
	"fmt"
	"regexp"
)

const (
	invalidErr = "This is an invalid error, please report it to the developer of this module, err: "
	findVerbs  = "(%(v|(#v)|T|t|[b-d]|o|0|q|x|X|U|[e-g]|[E-G]|s|p|(([0-9]|.)*f)))"
)

type AlmiErrorMsg string

type almiError struct {
	msg string
}

const (
	// errors for util
	SepAtoiErr            AlmiErrorMsg = "separator must be specified for AlmiAtoi func when 'val' is of type []T"
	SepStrErr             AlmiErrorMsg = "separator must be specified for AlmiStr func when 'val' is of type []T"
	SepAtobErr            AlmiErrorMsg = "separator must be specified for AlmiAtob func when 'val' is of type []T"
	SepAtoRBErr           AlmiErrorMsg = "separator must be specified for AlmiAtoRB func when 'val' is of type []T"
	AtoRBConversionFailed AlmiErrorMsg = "failed to convert string to int to convert to rune/byte"

	// config errors
	SepUndefErr                   AlmiErrorMsg = "Field: '%s': slice types must specify a separator character in their brackets"
	ConstraintUnknownErr          AlmiErrorMsg = "Constraint: '%s' at Field: '%s', is unknown to almi config"
	UnrecognizedTypeErr           AlmiErrorMsg = "AlmiConfig: unrecognized type: '%s'"
	FieldStructTagTypeMismatchErr AlmiErrorMsg = "Field: '%s' Type: '%s' in '%s' struct does not match the constraint Type: '%s' in '%s' struct tag"
	EnvConstraintUndefErr         AlmiErrorMsg = "'env=' constraint must be defined for all fields of the config, constraint not found for field: '%s'"
	FieldRequiredErr              AlmiErrorMsg = "Field: '%s', is required"
	FailedToConvertTypeErr        AlmiErrorMsg = "failed to convert type of '%s' to %s from string"
)

func (aem AlmiErrorMsg) Build(args ...any) *almiError {
	msg := string(aem)
	matches := regexp.MustCompile(findVerbs).FindAllString(msg, -1)
	if len(matches) != len(args) {
		msg = invalidErr + msg
	}

	return &almiError{
		msg: fmt.Sprintf(msg, args...),
	}
}

func (ae *almiError) String() string {
	return fmt.Sprintf("AlmiConfigError: %s", ae.msg)
}

func (ae *almiError) Error() string {
	return ae.String()
}
