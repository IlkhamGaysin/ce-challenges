package main

import (
	"math"
	"testing"
)

func TestDist(t *testing.T) {
	if res := dist(point{0, 0}, point{10000, 10000}); int(res) != 14142 {
		t.Errorf("failed: dist (0, 0) (10000, 10000) is 14142.135623731, got %f", res)
	}
	if res := dist(point{3, 7}, point{1, 5}); int(1000*res) != 2828 {
		t.Errorf("failed: dist (3, 7) (1, 5) is 2.828427125, got %f", res)
	}
	if res := dist(point{1, 9}, point{8, 3}); int(1000*res) != 9219 {
		t.Errorf("failed: dist (1, 9) (8, 3) is 9.219544457, got %f", res)
	}
}

func BenchmarkDist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dist(point{float64(i%11) * 1000, float64((i/11)%11) * 1000}, point{float64((i/121)%11) * 1000, float64((i/1331)%11) * 1000})
	}
}

type point struct {
	x, y float64
}

func dist(a, b point) float64 {
	x, y := a.x-b.x, a.y-b.y
	return math.Sqrt(x*x + y*y)
}
