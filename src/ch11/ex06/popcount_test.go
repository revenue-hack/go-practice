package ex06

import (
	"testing"
)

func benchPopCount(b *testing.B, in uint64) {
	for i := 0; i < b.N; i++ {
		PopCount(in)
	}
}

func benchPopCount24(b *testing.B, in uint64) {
	for i := 0; i < b.N; i++ {
		PopCount24(in)
	}
}

func benchPopCount25(b *testing.B, in uint64) {
	for i := 0; i < b.N; i++ {
		PopCount25(in)
	}
}

func BenchmarkPopCount_1(b *testing.B)   { benchPopCount(b, 0x1) }
func BenchmarkPopCount24_1(b *testing.B) { benchPopCount24(b, 0x1) }
func BenchmarkPopCount25_1(b *testing.B) { benchPopCount25(b, 0x1) }

func BenchmarkPopCount_65535(b *testing.B)   { benchPopCount(b, 0xffff) }
func BenchmarkPopCount24_65535(b *testing.B) { benchPopCount24(b, 0xffff) }
func BenchmarkPopCount25_65535(b *testing.B) { benchPopCount25(b, 0xffff) }

func BenchmarkPopCount_max(b *testing.B)   { benchPopCount(b, 0xffffffffffffffff) }
func BenchmarkPopCount24_max(b *testing.B) { benchPopCount24(b, 0xffffffffffffffff) }
func BenchmarkPopCount25_max(b *testing.B) { benchPopCount25(b, 0xffffffffffffffff) }

/*
goos: darwin
goarch: amd64
pkg: ch11/ex06
BenchmarkPopCount_1-4           2000000000               0.44 ns/op
BenchmarkPopCount24_1-4         10000000               144 ns/op
BenchmarkPopCount25_1-4         500000000                3.81 ns/op
BenchmarkPopCount_65535-4       2000000000               0.44 ns/op
BenchmarkPopCount24_65535-4     10000000               156 ns/op
BenchmarkPopCount25_65535-4     100000000               11.1 ns/op
BenchmarkPopCount_max-4         2000000000               0.52 ns/op
BenchmarkPopCount24_max-4       10000000               150 ns/op
BenchmarkPopCount25_max-4       100000000               10.5 ns/op
*/
