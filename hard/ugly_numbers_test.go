package main

import "testing"

func TestPow3(t *testing.T) {
	if res := pow3(39); res != 4052555153018976267 {
		t.Errorf("failed: 3^39 = 4052555153018976267, got %d", res)
	}
	if res := pow3(1); res != 3 {
		t.Errorf("failed: 3^1 = 3, got %d", res)
	}
	if res := pow3(0); res != 1 {
		t.Errorf("failed: 3^0 = 1, got %d", res)
	}
}

func BenchmarkPow3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow3(i % 40)
	}
}

func BenchmarkPow3Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow3Naive(i % 40)
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
