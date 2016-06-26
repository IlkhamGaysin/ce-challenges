package main

import "testing"

func TestHappy(t *testing.T) {
	for k, v := range map[uint]uint{
		1: 1, 22: 8, 8: 64, 64: 52, 52: 29, 29: 85,
		85: 89, 89: 145, 145: 42, 42: 20, 20: 4, 4: 16, 16: 37,
		37: 58, 58: 89, 7: 49, 49: 97, 97: 130, 130: 10, 10: 1} {
		if r := happy(k); r != v {
			t.Errorf("failed: happy %d is %d, got %d", k, v, r)
		}
	}
}

func BenchmarkHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		happy(uint(i))
	}
}

func happy(a uint) (ret uint) {
	for a > 0 {
		b := a % 10
		ret += b * b
		a /= 10
	}
	return ret
}
