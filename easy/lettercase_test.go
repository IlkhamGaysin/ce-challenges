package main

import "testing"

func TestLcase(t *testing.T) {
	if rl, ru := lcase("-"); rl != 0 || ru != 0 {
		t.Errorf("failed: lcase - is 0 0, got %f %f", rl, ru)
	}
	if rl, ru := lcase("thisTHIS"); rl != 50 || ru != 50 {
		t.Errorf("failed: lcase thisTHIS is 50 50, got %f %f", rl, ru)
	}
	if rl, ru := lcase("aBCD"); rl != 25 || ru != 75 {
		t.Errorf("failed: lcase aBCD is 25 75, got %f %f", rl, ru)
	}
}

func BenchmarkLcase(b *testing.B) {
	h := []byte{32, 65, 66, 67, 97, 98, 99}
	for i := 0; i < b.N; i++ {
		var s []byte
		for j := i; j > 0; j /= 7 {
			s = append(s, h[i%7])
		}
		lcase(string(s))
	}
}

func lcase(s string) (float64, float64) {
	var c, l float64
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			l++
			c++
		} else if i >= 'A' && i <= 'Z' {
			c++
		}
	}
	if c == 0 {
		return 0, 0
	}
	r := 100 * l / c
	return r, 100 - r
}
