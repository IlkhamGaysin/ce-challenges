package main

import "testing"

func TestRev(t *testing.T) {
	for k, v := range map[int]int{
		0: 0, 1: 1, 10: 1, 15: 51, 29: 92, 110: 11,
		1001: 1001, 9990: 999, 9999: 9999} {
		if r := rev(k); r != v {
			t.Errorf("failed: rev %d is %d, got %d",
				k, v, r)
		}
	}
}

func BenchmarkRev(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rev(i % 10000)
	}
}

func rev(a int) (r int) {
	for ; a > 0; a /= 10 {
		r = 10*r + a%10
	}
	return r
}
