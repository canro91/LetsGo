package reflection

import (
	"reflect"
)

func walk(x interface{}, f func(input string)) {
	value := reflect.ValueOf(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), f)
	}

	switch value.Kind() {
	case reflect.String:
		f(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			walkValue(field)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := value.Recv(); ok; v, ok = value.Recv() {
			walk(v.Interface(), f)
		}
	case reflect.Func:
		valFnResult := value.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), f)
		}
	case reflect.Ptr:
		walkValue(value.Elem())
	}
}
