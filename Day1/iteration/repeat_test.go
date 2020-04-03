package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){

	t.Run("Repeat_ByDefault_RepeatsFiveTimes", func(t *testing.T){
		expected := "aaaaa"
		actual := Repeat("a", 5)
	
		if expected != actual {
			t.Errorf("Expected %q but was %q", expected, actual)
		}
	})

	t.Run("Repeat_Count_RepeatsCountTimes", func(t *testing.T){
		expected := "aaa"
		actual := Repeat("a", 3)
	
		if expected != actual {
			t.Errorf("Expected %q but was %q", expected, actual)
		}
	})
}

func ExampleRepeat(){
	fmt.Println(Repeat("Hello", 2))
	// Output:
	// HelloHello
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}