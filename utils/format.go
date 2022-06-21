package utils

import (
	"reflect"
	"strings"
)

func FormatStructWithMap(value reflect.Value, msgMap *map[string]string) reflect.Value {
	for i := 0; i < value.NumField(); i++ {
		fieldVal := value.Field(i)
		fieldVal.Set(FormatValueWithMap(fieldVal, msgMap))
	}
	return value
}

func FormatStringWithMap(value string, msgMap *map[string]string) string {
	if !strings.Contains(value, "{{") {
		return value
	}
	for msgFlag, msg := range *msgMap {
		value = strings.Replace(value, msgFlag, msg, -1)
	}
	return value
}

func FormatSliceWithMap(value reflect.Value, msgMap *map[string]string) reflect.Value {
	for i := 0; i < value.Len(); i++ {
		sliceVal := value.Index(i)
		sliceVal.Set(FormatValueWithMap(sliceVal, msgMap))
	}
	return value
}

func FormatMapWithMap(value reflect.Value, msgMap *map[string]string) reflect.Value {
	for _, key := range value.MapKeys() {
		val := value.MapIndex(key)
		tmp := reflect.New(val.Type())
		tmp.Elem().Set(val)
		tmp = FormatValueWithMap(tmp.Elem(), msgMap)
		value.SetMapIndex(key, tmp)
	}
	return value
}

func FormatInterfaceWithMap(value reflect.Value, msgMap *map[string]string) reflect.Value {
	tmp := reflect.New(value.Type())
	tmp.Elem().Set(value)
	tmp = FormatValueWithMap(tmp.Elem(), msgMap)
	return tmp
}

func FormatValueWithMap(value reflect.Value, msgMap *map[string]string) reflect.Value {
	switch value.Kind() {
	case reflect.String:
		value.SetString(FormatStringWithMap(value.String(), msgMap))
	case reflect.Slice, reflect.Array:
		value.Set(FormatSliceWithMap(value, msgMap))
	case reflect.Struct:
		value.Set(FormatStructWithMap(value, msgMap))
	case reflect.Map:
		value.Set(FormatMapWithMap(value, msgMap))
	case reflect.Interface:
		value.Set(FormatInterfaceWithMap(value.Elem(), msgMap))
	}
	return value
}

// FormatAnyWithMap 传入指针
func FormatAnyWithMap(obj any, msgMap *map[string]string) any {
	if len(*msgMap) == 0 {
		return obj
	}
	ref := reflect.ValueOf(obj)
	ref.Elem().Set(FormatValueWithMap(ref.Elem(), msgMap))
	return obj
}
