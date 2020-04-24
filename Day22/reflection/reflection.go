package reflection

import (
	"reflect"
)

func walk(x interface{}, f func(input string)){
	value := reflect.ValueOf(x)
	field := value.Field(0)
	f(field.String())
}