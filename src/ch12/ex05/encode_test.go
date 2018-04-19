package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	jsonBytes, err := Marshal(strangelove)
	if err != nil {
		t.Error("json marshal error")
	}
	var movie Movie
	if err := json.Unmarshal(jsonBytes, &movie); err != nil {
		t.Error("json unmarchal error")
	}
	if !reflect.DeepEqual(movie, strangelove) {
		t.Error("test encode error")
	}
}
