package main

import (
	"strconv" // remove as soon as Go supports binary literals
	"testing"
)

func TestG2d(t *testing.T) {
	for k, v := range map[string]uint{
		"10": 3, "101": 6, "1111": 10,
		"1110": 11, "1100001": 65} {
		b, _ := strconv.ParseInt(k, 2, 0)
		if r := g2d(uint(b)); r != v {
			t.Errorf("failed: g2d %s is %d, got %d",
				k, v, r)
		}
	}
}

func BenchmarkG2d(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g2d(uint(i % 128))
	}
}

func g2d(a uint) uint {
	a ^= a >> 4
	a ^= a >> 2
	a ^= a >> 1
	return a
}
