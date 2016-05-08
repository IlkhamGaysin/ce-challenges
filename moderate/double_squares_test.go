package main

import (
	"math"
	"testing"
)

func TestSq(t *testing.T) {
	for i := 1; i < 2147483647; i = 16*i/10 + 1 {
		if res := int(math.Sqrt(float64(i * i))); res != i {
			t.Errorf("failed: math.Sqrt %d = %d, got %d", i*i, i, res)
		}
		if res := sqInt(i * i); res != i {
			t.Errorf("failed: sqInt %d = %d, got %d", i*i, i, res)
		}
	}
}

func BenchmarkSq(b *testing.B) {
	for i := 1; i < b.N; i++ {
		_ = int(math.Sqrt(float64(i)))
	}
}

func BenchmarkSqInt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		sqInt(i)
	}
}

func sqInt(a int) (r int) {
	for r = 1; r*r < a; r++ {
	}
	return r
}
