package almitypes

import "reflect"

type ConfigValue struct {
	Field       reflect.StructField
	Tag         string
	Constraints []string
	Value       reflect.Value
}

type ConfigConstraint struct {
	FieldName string

	Required bool
	EnvName  string
	Type     string

	SliceType bool
	Separator string
}
