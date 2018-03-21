package ex05

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, word := range []struct {
		s        string
		sep      string
		expected int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c:d", ":", 4},
	} {
		words := strings.Split(word.s, word.sep)
		if got := len(words); got != word.expected {
			t.Errorf("Split(%q, %q) returned %d words, want %d", word.s, word.sep, got, word.expected)
		}
	}
}
