package goat

import (
	"reflect"
	"strings"
)

type fieldBlueprint struct {
	FieldName string
	ParamName string
	Sample    any    // the zero value of the field type // INVESTIGATE it says unused write
	GetFrom   string // query, path or body
}

// func (b fieldBlueprint) String() string {
// 	return fmt.Sprintf("{ %s:%s [sample: %v] getFrom:%s }", b.FieldName, b.ParamName, b.Sample, b.GetFrom)
// }

func compileBlueprints(v any) []fieldBlueprint {
	blueprints := []fieldBlueprint{}

	rv := reflect.ValueOf(v)
	t := rv.Type()

	n := t.NumField()
	for i := 0; i < n; i++ {
		bp := fieldBlueprint{}
		field := t.Field(i)
		bp.FieldName = field.Name

		tag := field.Tag.Get("goat")
		tags := strings.Split(tag, ",")

		if len(tags) > 0 && tags[0] != "" {
			bp.ParamName = tags[0]
		} else {
			bp.ParamName = strings.ToLower(bp.FieldName)
		}

		if len(tags) > 1 {
			bp.GetFrom = tags[1]
		} else {
			if field.Type.Kind() == reflect.Struct {
				bp.GetFrom = "body"
			} else {
				bp.GetFrom = "query"
			}
		}

		if reflect.Zero(field.Type).CanInterface() {
			bp.Sample = reflect.Zero(field.Type).Interface()
		}

		blueprints = append(blueprints, bp)
	}

	return blueprints
}
