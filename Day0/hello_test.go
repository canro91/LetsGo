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

// %q encloses a string inside quotes

// There is no method overloading

// https://stackoverflow.com/a/17891297
// := declaration + assignment
// = assignment only

func TestHello(t *testing.T) {

    assertAreEqual := func(t *testing.T, expected, actual string){
        t.Helper()

        if expected != actual {
            t.Errorf("Expected %q but was %q", expected, actual)
        }
    }

    t.Run("Hello_Name_ReturnsHelloWithName", func(t *testing.T){
        expected := "Hello, Alice!"
        actual := Hello("Alice", "")

        assertAreEqual(t, expected, actual)
    })

    t.Run("Hello_EmptyString_ReturnsHelloWorld", func(t *testing.T){
        expected := "Hello, world!"
        actual := Hello("", "")

        assertAreEqual(t, expected, actual)
    })

    t.Run("Hello_SpanishAndName_ReturnsHelloInSpanish", func(t *testing.T){
        expected := "Hola, Alice!"
        actual := Hello("Alice", "Spanish")

        assertAreEqual(t, expected, actual)
    })

    t.Run("Hello_FrenchAndName_ReturnsHelloInSpanish", func(t *testing.T){
        expected := "Bonjour, Alice!"
        actual := Hello("Alice", "French")

        assertAreEqual(t, expected, actual)
    })
}


