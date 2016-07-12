package main

import "testing"

type tuple struct {
	x1, y1, x2, y2 int
}

func TestDist2(t *testing.T) {
	for k, v := range map[tuple]int{
		tuple{0, 0, 10, 10}: 200,
		tuple{3, 7, 1, 5}:   8,
		tuple{1, 9, 8, 3}:   85} {
		if r := dist2(point{k.x1, k.y1}, point{k.x2, k.y2}); r != v {
			t.Errorf("failed: dist2 (%d, %d) (%d, %d) is %d, got %d",
				k.x1, k.y1, k.x2, k.y2, v, r)
		}
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
