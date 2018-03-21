package ex02

import (
	"math/rand"
	"testing"
	"time"
)

const loopCount = 1000

func intSet_Add() {
	var intset IntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < loopCount; i++ {
		intset.Add(rand.Intn(10))
	}
}

func mapIntSet_Add() {
	var intset MapIntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < loopCount; i++ {
		intset.Add(rand.Intn(10))
	}
}

func BenchmarkIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSet_Add()
	}
}

func BenchmarkMapIntSet_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapIntSet_Add()
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSet_UnionWith()
	}
}
func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapIntSet_UnionWith()
	}
}

func intSet_UnionWith() {
	var x IntSet
	var y IntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < loopCount; i++ {
		x.Add(rand.Intn(10))
		y.Add(rand.Intn(10))
	}
	x.UnionWith(&y)
}

func mapIntSet_UnionWith() {
	var x MapIntSet
	var y MapIntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < loopCount; i++ {
		x.Add(rand.Intn(10))
		y.Add(rand.Intn(10))
	}
	x.UnionWith(&y)
}
