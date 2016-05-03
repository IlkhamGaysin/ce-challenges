package main

import (
	"strconv" // remove as soon as Go supports binary literals
	"testing"
)

func TestG2d(t *testing.T) {
	g := map[string]uint{"10": 3, "101": 6, "1111": 10,
		"1110": 11, "1100001": 65}
	for k, v := range g {
		b, _ := strconv.ParseInt(k, 2, 0)
		if res := g2d(uint(b)); res != v {
			t.Errorf("failed: g2d %s is %d, got %d", k, v, res)
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
