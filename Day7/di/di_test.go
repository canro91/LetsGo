package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alice")

	got := buffer.String()
	want := "Hello, Alice"

	if got != want {
		t.Errorf("Expected %s but was %s", want, got)
	}
}