package util

import (
	"encoding/json"
	"reflect"
	"strings"
)

func Bind(obj interface{}, data interface{}) error {
	if !containJsonTag(data) {
		return bindJson(obj, FieldToMap(data, "form"))
	}

	return bindJson(obj, data)
}

func bindJson(obj interface{}, data interface{}) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonStr, obj); err != nil {
		return err
	}

	return nil
}

func containJsonTag(obj interface{}) bool {
	if reflect.ValueOf(obj).Kind() == reflect.Slice {
		return true
	}

	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case IsStruct(objT):
	case IsStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return false
	}

	for i := 0; i < objT.NumField(); i++ {
		if objV.Field(i).Kind() == reflect.Struct {
			if containJsonTag(objV.Field(i).Interface()) {
				return true
			}
			continue
		}
		// just judgement the first tag is contain json
		tag := reflect.StructTag(objT.Field(i).Tag)
		if tag.Get("json") != "" {
			return true
		}
		return false
	}

	return false
}

func FieldToMap(in interface{}, tagFlag string) map[string]interface{} {
	out := make(map[string]interface{})

	inT := reflect.TypeOf(in)
	inV := reflect.ValueOf(in)
	switch {
	case IsStruct(inT):
	case IsStructPtr(inT):
		inT = inT.Elem()
		inV = inV.Elem()
	default:
		return nil
	}

	for i := 0; i < inT.NumField(); i++ {
		if inV.Field(i).Kind() == reflect.Struct {
			out = mergeMap(out, FieldToMap(inV.Field(i).Interface(), tagFlag))
			continue
		}
		var field string
		tag := reflect.StructTag(inT.Field(i).Tag)
		if tag.Get(tagFlag) != "" {
			field = tag.Get(tagFlag)
		}
		if field == "" && tag.Get("json") != "" {
			field = tag.Get("json")
		}
		if field == "" && tag.Get("form") != "" {
			field = tag.Get("form")
		}

		if inV.Field(i).IsZero() {
			continue
		}
		// compatible support default value situation
		field = extractFieldSpec(field)
		out[field] = inV.Field(i).Interface()
	}

	return out
}

func extractFieldSpec(field string) string {
	var spec string
	index := strings.Index(field, ",")
	spec = field
	if index != -1 {
		spec = field[:index]
	}

	return spec
}

func IsStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func IsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func mergeMap(dest, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		dest[k] = v
	}

	return dest
}
