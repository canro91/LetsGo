package reflection

import (
	"testing"
)

func TestWalk(t *testing.T){
	expected := "Alice"
	var actual []string

	x := struct{
		name string
	}{expected}

	walk(x, func(input string){
		actual = append(actual, input)
	})

	if len(actual) != 1 {
		t.Errorf("Expected %d, but was %d", 1, len(actual))
	}
	if actual[0] != expected {
		t.Errorf("Expected %s, but was %s", expected, actual[0])
	}
}