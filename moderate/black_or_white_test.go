package main

import "testing"

func TestKernighan(t *testing.T) {
	b := map[int]int{0: 0, 1: 1, 2: 1, 13: 3, 42: 3, 127: 7, 8192: 1,
		8200: 2, 4194305: 2}
	for k, v := range b {
		if res := kernighan(k); res != v {
			t.Errorf("failed: kernighan %d is %d, got %d", k, v, res)
		}
	}
}

func BenchmarkKernighan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kernighan(i)
	}
}

func kernighan(a int) (r int) {
	for ; a > 0; a &= a - 1 {
		r++
	}
	return r
}
