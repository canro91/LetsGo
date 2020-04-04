package arrays

import (
	"testing"
)

func TestSum(t *testing.T){
	array := [5]int{ 1, 2, 3, 4, 5}

	expected := 15
	sum := Sum(array)
	
	if expected != sum {
		t.Errorf("Expected %d but was %d", expected, sum)
	}
}