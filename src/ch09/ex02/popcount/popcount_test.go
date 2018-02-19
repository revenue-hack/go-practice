package popcount

import (
	"ch02/ex03/popcount"
	"testing"
)

func TestPopCount(t *testing.T) {
	if popcount.PopConut(10) != 2 {
		t.Error("TestPopCount func popCount not 10")
	}
	if popcount.PopConutOld(10) != 2 {
		t.Error("TestPopCount func popCountOld not 10")
	}
}
