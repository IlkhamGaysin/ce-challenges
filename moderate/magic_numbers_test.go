package main

import "testing"

func TestIsMagic(t *testing.T) {
	m := []int{1, 9, 35, 37, 174, 1267, 3562, 6712, 6392, 9263, 9627}
	n := []int{10, 12, 18, 175, 1624, 2715, 3261, 6372, 7216, 9876}
	for _, i := range m {
		if !isMagic(i) {
			t.Errorf("failed: isMagic %d is true, got false", i)
		}
	}
	for _, i := range n {
		if isMagic(i) {
			t.Errorf("failed: isMagic %d is false, got true", i)
		}
	}
}

func BenchmarkIsMagic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMagic(i%10000 + 1)
	}
}

func isMagic(a int) bool {
	var (
		dig, r uint
		ns     []uint
	)
	for a > 0 {
		r = uint(a % 10)
		if r == 0 || dig&(1<<r) > 0 {
			return false
		}
		dig |= 1 << r
		ns = append(ns, r)
		a /= 10
	}
	dig, r = 0, 0
	for _ = range ns {
		r = (r + ns[(uint(len(ns))-1-r)]) % uint(len(ns))
		if dig&(1<<r) > 0 {
			return false
		}
		dig |= 1 << r
	}
	return r == 0
}
