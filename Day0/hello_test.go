package main

import "testing"

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

    t.Run("Hello_SpanishCodeAndName_ReturnsHelloInSpanish", func(t *testing.T){
        expected := "Hola, Alice!"
        actual := Hello("Alice", "es")

        assertAreEqual(t, expected, actual)
    })

    t.Run("Hello_FrenchAndName_ReturnsHelloInFrench", func(t *testing.T){
        expected := "Bonjour, Alice!"
        actual := Hello("Alice", "French")

        assertAreEqual(t, expected, actual)
    })

    t.Run("Hello_FrenchCodeAndName_ReturnsHelloInFrench", func(t *testing.T){
        expected := "Bonjour, Alice!"
        actual := Hello("Alice", "fr")

        assertAreEqual(t, expected, actual)
    })
}
