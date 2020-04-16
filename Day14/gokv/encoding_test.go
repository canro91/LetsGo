package gokv

import (
	"bytes"
	"encoding/gob"
	"testing"
)

type Square struct {
	Length int
}

func TestEncoding(t *testing.T) {
	square := Square{Length: 2}

	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(&square)
	buffer.Bytes()

	decoder := gob.NewDecoder(&buffer)
	var decodedSquare Square
	decoder.Decode(&decodedSquare)

	if decodedSquare != square {
		t.Errorf("Expected %v but was %v", square, decodedSquare)
	}
}
