package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCardNumber(t *testing.T) {
	h := map[string]bool{
		"6011594003199511":    false,
		"5537021367976815":    true,
		"5574836380229735":    false,
		"30448507939130":      false,
		"6370167590346211774": true}
	for k, v := range h {
		if res := cardNumber(k); res != v {
			t.Errorf("failed: cardNumber %s is %t, got %t", k, v, res)
		}
	}
}

func BenchmarkCardNumber(b *testing.B) {
	for i := 1; i < b.N; i++ {
		cardNumber(fmt.Sprint(uint64(i) + 1000000000000000))
	}
}

func cardNumber(q string) bool {
	s := strings.Split(q, "")
	t := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		fmt.Sscan(s[i], &t[i])
	}
	for i := len(t) - 2; i >= 0; i -= 2 {
		t[i] *= 2
		if t[i] > 9 {
			t[i] = t[i]%10 + 1
		}
	}
	var su int
	for i := 0; i < len(t); i++ {
		su += t[i]
	}
	return su%10 == 0
}
