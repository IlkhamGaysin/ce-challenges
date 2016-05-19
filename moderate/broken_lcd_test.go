package main

import "testing"

func TestValid(t *testing.T) {
	if valid(1, true, 2) {
		t.Error("failed: valid 1 true 2 is false, got true")
	}
	if !valid(3, true, 7) {
		t.Error("failed: valid 7 true 3 is true, got false")
	}
	if valid(3, false, 6) {
		t.Error("failed: valid 6 false 3 is false, got true")
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valid(i%256, (i/256)&1 == 1, (i/512)%256)
	}
}

func valid(y int, d bool, x int) bool {
	if d && (x&1 == 0) {
		return false
	}
	return y&x == y
}
