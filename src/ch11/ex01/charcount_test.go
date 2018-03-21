package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	var readers = []struct {
		str     []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			[]byte("hoge"),
			map[rune]int{'h': 1, 'o': 1, 'g': 1, 'e': 1},
			[]int{0, 4, 0, 0, 0},
			0,
		},
		{
			[]byte("ほげ"),
			map[rune]int{'ほ': 1, 'げ': 1},
			[]int{0, 0, 0, 2, 0},
			0,
		},
		{
			[]byte("あああ, aaa"),
			map[rune]int{'あ': 3, 'a': 3, ',': 1, ' ': 1},
			[]int{0, 5, 0, 3, 0},
			0,
		},
		{
			[]byte("あiaあ, aa　a200"),
			map[rune]int{'あ': 2, 'a': 4, ',': 1, ' ': 1, '2': 1, '0': 2, 'i': 1, '　': 1},
			[]int{0, 10, 0, 3, 0},
			0,
		},
		{
			[]byte("Google\200"),
			map[rune]int{'G': 1, 'o': 2, 'g': 1, 'l': 1, 'e': 1},
			[]int{0, 6, 0, 0, 0},
			1,
		},
	}
	for _, reader := range readers {
		counts, utflen, invalid := charCount(bytes.NewReader(reader.str))
		for r, count := range counts {
			if count != reader.counts[r] {
				t.Errorf("reder is %v counts is %v and rune is %v want is %v\n", string(reader.str), count, string(r), reader.counts[r])
			}
		}
		for i, utf := range utflen {
			if utf != reader.utflen[i] {
				t.Errorf("index is %d reder is %v utflen is %d want is %d\n", i, string(reader.str), utf, reader.utflen[i])
			}
		}
		if invalid != reader.invalid {
			t.Errorf("reder is %v invalid is %d want is %d\n", string(reader.str), invalid, reader.invalid)
		}

	}
}
