package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T){

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Single property",
			struct{
				name string
			}{"Alice"},
			[]string{"Alice"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T){
			var actual []string

			walk(test.Input, func(input string){
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("Expected %s, but was %s", test.ExpectedCalls, actual)
			}
		})
	}
}