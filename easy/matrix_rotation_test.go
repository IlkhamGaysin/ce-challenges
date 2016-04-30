package main

import (
	"fmt"
	"testing"
)

func TestSq(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if sq(i*i) != i {
			t.Error("failed: sq " + fmt.Sprint(i*i))
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
