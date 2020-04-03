package iteration

import "testing"

func TestRepeat(t *testing.T){
	expected := "aaaaa"
	actual := Repeat("a")

	if expected != actual {
		t.Errorf("Expected %q but was %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}