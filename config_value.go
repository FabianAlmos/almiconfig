package almi

import (
	"almi/lexer"
	"reflect"
)

type configValue struct {
	Field       reflect.StructField
	Tag         string
	Constraints []string
	Value       reflect.Value
}

func getConfigValueByIndex(cfg reflect.Value, i int) *configValue {
	field := cfg.Type().Field(i)
	tag := field.Tag.Get(almi)

	lxr := lexer.NewLexer(tag)

	return &configValue{
		Field:       field,
		Tag:         tag,
		Constraints: lxr.Tokenize(),
		Value:       cfg.FieldByIndex([]int{i}),
	}
}
