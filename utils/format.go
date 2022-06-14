package utils

import (
	"reflect"
	"strings"
)

func FormatStructWithMap(value reflect.Value, msgMap map[string]string) reflect.Value {
	for i := 0; i < value.NumField(); i++ {
		fieldVal := value.Field(i)
		fieldVal.Set(FormatValueWithMap(fieldVal, msgMap))
	}
	return value
}

func FormatStringWithMap(str string, msgMap map[string]string) string {
	if !strings.Contains(str, "{{") {
		return str
	}
	for msgFlag, msg := range msgMap {
		str = strings.Replace(str, msgFlag, msg, -1)
	}
	return str
}

func FormatSliceWithMap(value reflect.Value, msgMap map[string]string) reflect.Value {
	for i := 0; i < value.Len(); i++ {
		sliceVal := value.Index(i)
		sliceVal.Set(FormatValueWithMap(sliceVal, msgMap))
	}
	return value
}

func FormatValueWithMap(value reflect.Value, msgMap map[string]string) reflect.Value {
	switch value.Kind() {
	case reflect.String:
		value.SetString(FormatStringWithMap(value.String(), msgMap))
	case reflect.Slice, reflect.Array:
		value.Set(FormatSliceWithMap(value, msgMap))
	case reflect.Struct:
		value.Set(FormatStructWithMap(value, msgMap))
	}
	return value
}

// FormatAnyWithMap 传入指针
func FormatAnyWithMap(obj any, msgMap map[string]string) any {
	if len(msgMap) == 0 {
		return obj
	}
	ref := reflect.ValueOf(obj)
	ref.Elem().Set(FormatValueWithMap(ref.Elem(), msgMap))
	return obj
}
