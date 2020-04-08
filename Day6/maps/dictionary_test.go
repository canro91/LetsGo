package maps

import (
	"testing"
)

func TestDictionary(t *testing.T){
	t.Run("Search_ExistingKey_ReturnsValue", func(t *testing.T){
		dictionary := Dictionary{"test": "This is a test"}

		expected := "This is a test"
		value, _ := dictionary.Search("test")
	
		assertAreEqual(t, expected, value)	
	})

	t.Run("Search_UnknownKey_ReturnsError", func(t *testing.T){
		dictionary := Dictionary{"test": "This is a test"}

		_, error := dictionary.Search("unknown")

		assertError(t, error, ErrKeyNotFound)
	})

	t.Run("Add_Key_AddsKey", func(t *testing.T){
		dictionary := Dictionary{}
		dictionary.Add("test", "This is a test")

		assertContains(t, dictionary, "test", "This is a test")
	})

	t.Run("Add_ExistingKey_ReturnsError", func(t *testing.T){
		dictionary := Dictionary{"test": "This is a test"}

		error := dictionary.Add("test", "This is another test")

		assertError(t, error, ErrKeyAlreadyAddedd)
		assertContains(t, dictionary, "test", "This is a test")
	})

	t.Run("Update_ExistingKey_ReturnsError", func(t *testing.T){
		dictionary := Dictionary{"test": "This is a test"}

		dictionary.Update("test", "This is another test")

		assertContains(t, dictionary, "test", "This is another test")
	})
	
	t.Run("Update_NotExistingKey_ReturnsError", func(t *testing.T){
		dictionary := Dictionary{}

		error := dictionary.Update("test", "This is another test")

		assertError(t, error, ErrKeyNotFound)
	})

	t.Run("Delete_ExistingKey_DeletesKey", func(t *testing.T){
		dictionary := Dictionary{"test": "This is a test"}

		dictionary.Delete("test")

		_, err := dictionary.Search("test")

		assertError(t, ErrKeyNotFound, err)
	})
}

func assertAreEqual(t *testing.T, expected, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected %q but was %q", expected, actual)
	}
}

func assertError(t *testing.T, expected, actual error){
	t.Helper()

	if expected == nil {
		t.Fatal("Expected error but was nil")	
	}

	if expected != actual {
		t.Errorf("Error message expected to contain %q but was %q", expected, actual)
	}
}

func assertNoError(t *testing.T, err error){
	t.Helper()

	if err != nil {
		t.Errorf("Error not expected error but was %q", err)	
	}
}

func assertContains(t *testing.T, d Dictionary, key, expected string) {
	value, error := d.Search(key)

	assertAreEqual(t, expected, value)
	assertNoError(t, error)
}