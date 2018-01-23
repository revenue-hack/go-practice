package main

import "testing"

var cases = []struct {
	wordExpected WordCounter
	lineExpected LineCounter
	bytes        []byte
}{
	{7, 3, []byte("あいうえお かかかかか hohoho　あぁぁｌ\nhoge hgoe\naaaa")},
	{4, 2, []byte("あいうえお かかかかか hohoho　あぁぁｌ\n   ")},
	{0, 4, []byte("\n\n\n\n")},
	{0, 0, []byte("")},
}

func TestLineCounter_Write(t *testing.T) {
	for _, c := range cases {
		var l LineCounter
		l.Write(c.bytes)
		if l != c.lineExpected {
			t.Errorf("don't expected value %d\t%v\n", l, c.lineExpected)
		}
	}
}

func TestWordCounter_Write(t *testing.T) {
	for _, c := range cases {
		var w WordCounter
		w.Write(c.bytes)
		if w != c.wordExpected {
			t.Errorf("don't expected value %d\t%v\n", w, c.lineExpected)
		}
	}
}
