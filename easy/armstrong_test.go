package main

import "testing"

func TestPowi(t *testing.T) {
	if res := powi(9, 19); res != 1350851717672992089 {
		t.Errorf("failed: 9^19 = 1350851717672992089, got %d", res)
	}
	if res := powi(0, 8); res != 0 {
		t.Errorf("failed: 0^8 = 0, got %d", res)
	}
	if res := powi(0, 0); res != 1 {
		t.Errorf("failed: 0^0 = 1, got %d", res)
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
