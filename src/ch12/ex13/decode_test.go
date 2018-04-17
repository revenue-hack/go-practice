package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	data, err := Marshal(testMovie)
	if err != nil {
		t.Error("marshal error")
	}
	var expected Movie
	err = Unmarshal(data, &expected)
	if err != nil {
		t.Error("unmarshal error")
	}
	if !reflect.DeepEqual(expected, testMovie) {
		t.Error("test unmarshal error")
	}
}

func TestDecoder_Decode(t *testing.T) {
	data, err := Marshal(testMovie)
	if err != nil {
		t.Error("marshal error")
	}
	var expected Movie
	reader := bytes.NewReader(data)
	dec := &Decoder{reader}

	if err := dec.Decode(&expected); err != nil {
		t.Error("decode error")
	}
	if !reflect.DeepEqual(expected, testMovie) {
		t.Error("test decode error")
	}
}
