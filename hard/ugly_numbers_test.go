package main

import "testing"

func TestPow3(t *testing.T) {
	if pow3(39) != 4052555153018976267 {
		t.Error("failed: 3^39")
	}
	if pow3(1) != 3 {
		t.Error("failed: 3^1")
	}
	if pow3(0) != 1 {
		t.Error("failed: 3^0")
	}
}

func BenchmarkPow3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow3(i%40)
	}
}

func BenchmarkPow3Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow3Naive(i%40)
	}
}

func pow3(x int) int {
	y, ret := 3, 1
	for x > 0 {
		if x&1 > 0 {
			ret *= y
		}
		y *= y
		x >>= 1
	}
	return ret
}

func pow3Naive(a int) int {
	ret := 1
	for i := 0; i < a; i++ {
		ret *= 3
	}
	return ret
}
