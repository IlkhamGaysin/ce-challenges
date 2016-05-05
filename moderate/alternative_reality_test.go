package main

import "testing"

func TestAlter(t *testing.T) {
	r := map[int]int{4: 1, 17: 6, 100: 292}
	for k, v := range r {
		if res := alter(k, 50); res != v {
			t.Errorf("failed: alter %d, 50 is %d, got %d", k, v, res)
		}
	}
}

func BenchmarkAlter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alter(i%100+1, 50)
	}
}

var m map[int]int

func init() {
	m = map[int]int{50: 25, 25: 10, 10: 5, 5: 1}
}

func alter(n, c int) int {
	if n == 0 || c == 1 {
		return 1
	}
	if c > n {
		return alter(n, m[c])
	}
	return alter(n-c, c) + alter(n, m[c])
}
