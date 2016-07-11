package main

import "testing"

type tuple struct {
	y, x int
}

func TestPowi(t *testing.T) {
	for k, v := range map[tuple]int{
		tuple{9, 19}: 1350851717672992089,
		tuple{0, 8}:  0,
		tuple{0, 0}:  1} {
		if r := powi(k.y, k.x); r != v {
			t.Errorf("failed: %d^%d = %d, got %d",
				k.y, k.x, v, r)
		}
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
