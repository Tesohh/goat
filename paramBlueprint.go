package goat

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type fieldBlueprint struct {
	FieldName string
	ParamName string
	// Sample    any    // the zero value of the field type // INVESTIGATE it says unused write
	Type    reflect.Type
	GetFrom string // query, path or body
}

// func (b fieldBlueprint) String() string {
// 	return fmt.Sprintf("{ %s:%s [sample: %v] getFrom:%s }", b.FieldName, b.ParamName, b.Sample, b.GetFrom)
// }

func (b fieldBlueprint) CastInt(s string) (int, error) {
	if b.Type.Kind() == reflect.Int {
		return strconv.Atoi(s)
	} else {
		return 0, fmt.Errorf("can't cast to int a %s", b.Type.Kind())
	}
}

func (b fieldBlueprint) CastFloat64(s string) (float64, error) {
	if b.Type.Kind() == reflect.Float64 {
		return strconv.ParseFloat(s, 64)
	} else {
		return 0, fmt.Errorf("can't cast to float64 a %s", b.Type.Kind())
	}
}

func (b fieldBlueprint) CastFloat32(s string) (float32, error) {
	if b.Type.Kind() == reflect.Float32 {
		value, err := strconv.ParseFloat(s, 32)
		return float32(value), err
	} else {
		return 0, fmt.Errorf("can't cast to float32 a %s", b.Type.Kind())
	}
}

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

		// if reflect.Zero(field.Type).CanInterface() {
		// 	bp.Sample = reflect.Zero(field.Type).Interface()
		// }
		bp.Type = field.Type

		blueprints = append(blueprints, bp)
	}

	return blueprints
}
