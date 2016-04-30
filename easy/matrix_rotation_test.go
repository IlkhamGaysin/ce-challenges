package main

import "testing"

func TestSq(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if res := sq(i * i); res != i {
			t.Errorf("failed: sq %d = %d, got %d", i*i, i, res)
		}
	}
}

func BenchmarkSq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sq(i%10 + 1)
	}
}

func sq(a int) (r int) {
	for r = 1; r*r < a; r++ {
	}
	return r
}
