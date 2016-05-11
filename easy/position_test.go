package main

import "testing"

func TestXz(t *testing.T) {
	if !bitEqual(86, 2, 3) {
		t.Error("failed: bitEqual 86 2 3 should be true")
	}
	if bitEqual(125, 1, 2) {
		t.Error("failed: bitEqual 125 1 2 should be false")
	}
	if !bitEqual(1, 1, 1) {
		t.Error("failed: bitEqual 1 1 1 should be true")
	}
	if !bitEqual(0, 4, 16) {
		t.Error("failed: bitEqual 0 4 16 should be true")
	}
}

func BenchmarkBitEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitEqual(uint(i/256), uint((i/16)%16), uint(i%16))
	}
}

func bitEqual(n, p, q uint) bool {
	return ((n & (1 << (p - 1))) == 0) == ((n & (1 << (q - 1))) == 0)
}
