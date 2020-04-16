package gokv

import (
	"testing"
)

type Session struct {
	Code string
}

func TestKeyValueStore(t *testing.T) {
	store, err := Open("sessions.db")
	assertNoError(t, err)

	session := Session{"ABC123"}
	err = store.Put("session-1", session)
	assertNoError(t, err)

	var found Session
	err = store.Get("session-1", &found)
	assertNoError(t, err)
	if found != session {
		t.Errorf("Expected %v but was %v", session, found)
	}

	err = store.Delete("session-1")
	assertNoError(t, err)

	err = store.Get("session-1", &found)
	assertError(t, err, ErrKeyNotFound)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Error not expected error but was %q", err)
	}
}

func assertError(t *testing.T, expected, actual error) {
	t.Helper()

	if expected == nil {
		t.Fatal("Expected error but was nil")
	}

	if expected != actual {
		t.Errorf("Error message expected to contain %q but was %q", expected, actual)
	}
}
