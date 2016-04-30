package main

import "testing"

func TestPowi(t *testing.T) {
	if powi(9, 19) != 1350851717672992089 {
		t.Error("failed: 9^19")
	}
	if powi(0, 8) != 0 {
		t.Error("failed: 0^8")
	}
	if powi(0, 0) != 1 {
		t.Error("failed: 0^0")
	}
}

func BenchmarkPowi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powi(i%10, i%20)
	}
}

func BenchmarkPowiNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powiNaive(i%10, i%20)
	}
}

func powi(y, x int) int {
	ret := 1
	for x > 0 {
		if x&1 > 0 {
			ret *= y
		}
		y *= y
		x >>= 1
	}
	return ret
}

func powiNaive(a, b int) (ret int) {
	ret = 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}
