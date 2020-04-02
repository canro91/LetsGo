package main

import "fmt"
import "testing"

// Tests should be in a file XXX_test.go
// Test names should start with Test

// There isn't anything like Assert.AreEqual
func TestHello(t *testing.T) {
    expected := "Hello, world!"
    actual := Hello()

    if expected != actual {
        // %q encloses a string inside quotes
        t.Errorf("Expected %q but was %q", expected, actual)
    }
}

// Examples should contain an "Output" comment
func ExampleHello() {
    fmt.Println(Hello())
    // Output:
    // Hello, world!
}
