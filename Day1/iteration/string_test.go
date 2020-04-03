package iteration

import (
	"strings"
	"testing"
)

func TestString(t *testing.T){

	assertIsTrue := func(t *testing.T, v bool){
		t.Helper()

		if !v {
			t.Errorf("Expected true, but was: <%v>", v)
		}
	}
	
	assertIsFalse := func(t *testing.T, v bool){
		t.Helper()

		if v {
			t.Errorf("Expected false, but was: <%v>", v)
		}
	}

	assertAreEqual := func(t *testing.T, expected, actual string){
        t.Helper()

        if expected != actual {
            t.Errorf("Expected %q but was %q", expected, actual)
        }
    }

	t.Run("Contains", func(t *testing.T){
		assertIsTrue(t, strings.Contains("Hello, World", "World"))
		assertIsFalse(t, strings.Contains("Hello, World", "hola"))

		assertIsFalse(t, strings.Contains("Hello, World", "world"))
		assertIsFalse(t, strings.Contains("Hello, World", "wOrLd"))
	})

	t.Run("Index", func(t *testing.T){
		assertIsTrue(t, strings.Index("Hello, World", "World") == 7)

		assertIsTrue(t, strings.Index("Hello, World", "hola") == -1)
	})

	t.Run("ReplaceAll", func(t *testing.T){
		assertAreEqual(t, "Helli, Wirld", strings.ReplaceAll("Hello, World", "o", "i"))

		assertAreEqual(t, "Hello, World", strings.ReplaceAll("Hello, World", "abc", "i"))
	})

	t.Run("Title vs ToTitle", func(t *testing.T){
		assertAreEqual(t, "Hello, World", strings.Title("hello, world"))

		assertAreEqual(t, "HELLO, WORLD", strings.ToTitle("Hello, World"))

		assertAreEqual(t, strings.ToTitle("Hello, World"), strings.ToUpper("Hello, World"))
	})

	t.Run("Trim", func(t *testing.T){
		assertAreEqual(t, "Hello, World", strings.Trim("Hello, World!", "!"))

		assertAreEqual(t, ", World", strings.TrimPrefix("Hello, World", "Hello"))

		assertAreEqual(t, "Hello, World", strings.TrimSpace("   Hello, World   "))
	})
}