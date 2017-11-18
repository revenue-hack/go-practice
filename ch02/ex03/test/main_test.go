package main

import (
	"ch02/ex03/popcount"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		popcount.PopConut(uint64(i))
	}
}

func BenchmarkOld(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		popcount.PopConutOld(uint64(i))
	}
}

func TestPopCount(t *testing.T) {
	if popcount.PopConut(10) != 2 {
		t.Error("TestPopCount func popCount not 10")
	}
	if popcount.PopConutOld(10) != 2 {
		t.Error("TestPopCount func popCountOld not 10")
	}
}
