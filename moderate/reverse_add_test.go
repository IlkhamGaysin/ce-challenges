package main

import "testing"

func TestRev(t *testing.T) {
	r := map[int]int{0: 0, 1: 1, 10: 1, 15: 51, 29: 92, 110: 11,
		1001: 1001, 9990: 999, 9999: 9999}
	for k, v := range r {
		if res := rev(k); res != v {
			t.Errorf("failed: rev %d is %d, got %d", k, v, res)
		}
	}
}

func BenchmarkRev(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rev(i % 10000)
	}
}

func rev(a int) (ret int) {
	for ; a > 0; a /= 10 {
		ret = 10*ret + a%10
	}
	return ret
}
