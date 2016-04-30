package main

import "testing"

func TestXz(t *testing.T) {
	if !xz(1, 5) {
		t.Error("failed: 1, 5")
	}
	if xz(1, 7) {
		t.Error("failed: 1, 7")
	}
	if !xz(7, 10922) {
		t.Error("failed: 7, 10922")
	}
}

func BenchmarkXz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xz(i%11, (i/11)%995+5)
	}
}

func xz(n, a int) bool {
	for a > 0 && n >= 0 {
		if a&1 == 0 {
			n--
		}
		a >>= 1
	}
	return n == 0
}
