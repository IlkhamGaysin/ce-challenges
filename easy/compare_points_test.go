package main

import "testing"

func TestComparePoints(t *testing.T) {
	if res := comparePoints(0, 0, 1, 5); res != "NE" {
		t.Errorf("failed: comparePoints(0, 0, 1, 5) is NE, got %s", res)
	}
	if res := comparePoints(12, 13, 12, 13); res != "here" {
		t.Errorf("failed: comparePoints(12, 13, 12, 13) is here, got %s", res)
	}
	if res := comparePoints(0, 1, 0, 5); res != "N" {
		t.Errorf("failed: comparePoints(0, 1, 0, 5) is N, got %s", res)
	}
	if res := comparePoints(0, 0, 1, -5); res != "SE" {
		t.Errorf("failed: comparePoints(0, 0, 1, -5) is SE, got %s", res)
	}
}

func BenchmarkComparePoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x1 := i%21 - 10
		y1 := (i/21)%21 - 10
		x2 := (i/441)%21 - 10
		y2 := (i/9261)%21 - 10
		_ = comparePoints(x1, y1, x2, y2)
	}
}

func comparePoints(x1, y1, x2, y2 int) string {
	switch {
	case x1 == x2:
		if y1 == y2 {
			return "here"
		} else if y1 < y2 {
			return "N"
		} else {
			return "S"
		}
	case y1 == y2:
		if x1 < x2 {
			return "E"
		} else {
			return "W"
		}
	case x1 < x2:
		if y1 < y2 {
			return "NE"
		} else {
			return "SE"
		}
	default:
		if y1 < y2 {
			return "NW"
		} else {
			return "SW"
		}
	}
}
