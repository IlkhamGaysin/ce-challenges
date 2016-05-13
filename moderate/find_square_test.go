package main

import "testing"

func TestDist2(t *testing.T) {
	if res := dist2(point{0, 0}, point{10, 10}); res != 200 {
		t.Errorf("failed: dist2 (0, 0) (10, 10) is 200, got %d", res)
	}
	if res := dist2(point{3, 7}, point{1, 5}); res != 8 {
		t.Errorf("failed: dist2 (3, 7) (1, 5) is 8, got %d", res)
	}
	if res := dist2(point{1, 9}, point{8, 3}); res != 85 {
		t.Errorf("failed: dist2 (1, 9) (8, 3) is 85, got %d", res)
	}
}

func BenchmarkDist2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dist2(point{i % 11, (i / 11) % 11}, point{(i / 121) % 11, (i / 1331) % 11})
	}
}

type point struct {
	x, y int
}

func dist2(a, b point) int {
	x, y := a.x-b.x, a.y-b.y
	return x*x + y*y
}
