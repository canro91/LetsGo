package main

import "testing"

// Tests should be in a file XXX_test.go
// Test names should start with Test
// Tests should have a single param *testing.T
// There isn't anything like Assert.AreEqual by default

// Examples should contain an "Output" comment
//func ExampleHello() {
//    fmt.Println(Hello())
//    // Output:
//    // Hello, world!
//}

// Param types follow param names in functions
// Return type are between closing parenthesis and opening bracket

// Semicolons are optional at the end of line

// Unused imports give a compilation error

func TestHello(t *testing.T) {
    expected := "Hello, Alice!"
    actual := Hello("Alice")

    if expected != actual {
        // %q encloses a string inside quotes
        t.Errorf("Expected %q but was %q", expected, actual)
    }
}

