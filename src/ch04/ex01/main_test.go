package main

import (
	"testing"
	"crypto/sha256"
)

func TestCompareSha256OfNormal(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	if compareSha256(c1, c2) != 125 {
		t.Error("Test Normal Error")
	}
}

func TestCompareSha256OfSameNormal(t *testing.T) {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("X"))
	if compareSha256(c1, c2) != 0 {
		t.Error("Test SameNormal Error")
	}
}
