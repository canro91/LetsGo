package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){

	t.Run("Sum_ArrayOfAnyLength_ReturnsSum", func (t *testing.T) {
		array := []int{ 1, 2, 3 }

		expected := 6
		sum := Sum(array)
		
		if expected != sum {
			t.Errorf("Expected %d but was %d", expected, sum)
		}
	})
}

func TestSumAll(t *testing.T){
	expected := []int{ 6, 3 }
	sum := SumAll([]int{ 1,2,3 }, []int { 1,2 })
	
	if !reflect.DeepEqual(expected, sum) {
		t.Errorf("Expected %d but was %d", expected, sum)
	}	
}

func TestSumTails(t *testing.T){
	checkSums := func(t *testing.T, expected, actual []int){
		t.Helper()

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %d but was %d", expected, actual)
		}
	}


	t.Run("SumTails_NonEmptyArrays_ReturnsSum", func(t *testing.T){
		expected := []int{ 5, 2 }
		sum := SumTails([]int{ 1,2,3 }, []int { 1,2 })
		
		checkSums(t, expected, sum)
	})

	t.Run("SumTails_EmptyArray_ReturnsSum", func(t *testing.T){
		expected := []int{ 0, 2 }
		sum := SumTails([]int{}, []int { 1,2 })
		
		checkSums(t, expected, sum)
	})
}