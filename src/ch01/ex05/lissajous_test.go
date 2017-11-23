package main

import (
	"testing"
)

func TestCalcColorIndex(t *testing.T) {
	x, y, color := calcColorIndex(1.0, 1.0, 100)
	if x != 200 {
		t.Error("TestCalcColorIndex func x error")
	}
	if y != 200 {
		t.Error("TestCalcColorIndex func y error")
	}
	if color != 1 {
		t.Error("TestCalcColorIndex func color error")
	}
}

func TestCalcXY(t *testing.T) {
	x, y := calcXY(1.0, 1.0, 1.0)
	if x != 0.8414709848078965 {
		t.Error("TestCalcXY func x error")
	}
	if y != 0.9092974268256816 {
		t.Error("TestCalcXY func y error")
	}
}

func TestCalcRect(t *testing.T) {
	v, w, x, y := calcRect(100)
	if v != 0 {
		t.Error("TestCalcRect func v error")
	}
	if w != 0 {
		t.Error("TestCalcRect func w error")
	}
	if x != 201 {
		t.Error("TestCalcRect func x error")
	}
	if y != 201 {
		t.Error("TestCalcRect func y error")
	}
}
