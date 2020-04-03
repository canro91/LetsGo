package integers

import (
    "testing"
    "fmt"
)

func TestAdder(t *testing.T){
    expected := 4
    actual := Add(2, 2)

    if expected != actual {
        t.Errorf("Expected '%d' but was '%d'", expected, actual)
    }
}

func ExampleAdd(){
    fmt.Println(Add(2, 2))
    // Output:
    // 4
}
