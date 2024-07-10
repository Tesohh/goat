package goat

import (
	"fmt"
	"net/http"
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

func (b fieldBlueprint) Cast(s string) (any, error) {
	switch b.Type.Kind() {
	case reflect.Int:
		return strconv.Atoi(s)
	case reflect.Float64:
		return strconv.ParseFloat(s, 64)
	case reflect.Float32:
		value, err := strconv.ParseFloat(s, 32)
		return float32(value), err
	case reflect.String:
		return s, nil
	}

	// TODO: go has a reflect.Convert function

	return reflect.Zero(b.Type).Interface(), fmt.Errorf("cannot cast from string to %s", b.Type.Kind())
}

func (b fieldBlueprint) SetField(params reflect.Value, s *Server, r *http.Request) error {
	field := params.FieldByName(b.FieldName)

	if !field.IsValid() {
		return fmt.Errorf("field %s is invalid", b.FieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set params field %s", b.FieldName)
	}

	if b.GetFrom == "query" {
		rawQuery := r.URL.Query().Get(b.ParamName)

		castedValue, err := b.Cast(rawQuery)
		if err != nil {
			return err
		}

		field.Set(reflect.ValueOf(castedValue))
	} else if b.GetFrom == "path" {
		panic("not implemented")
	} else if b.GetFrom == "body" {
		panic("not implemented")
	} else {
		fmt.Println("Unknown GetFrom option", b.GetFrom)
	}

	return nil
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
