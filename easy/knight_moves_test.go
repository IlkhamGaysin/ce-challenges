package main

import "testing"

func TestPos(t *testing.T) {
	if res := pos(0, 0); res != "a1" {
		t.Error("failed: pos 0 0 is a1, got %s", res)
	}
	if res := pos(7, 7); res != "h8" {
		t.Error("failed: pos 7 7 is h8, got %s", res)
	}
}

func TestPosConcat(t *testing.T) {
	if res := posConcat(0, 0); res != "a1" {
		t.Error("failed: posConcat 0 0 is a1, got %s", res)
	}
	if res := posConcat(7, 7); res != "h8" {
		t.Error("failed: posConcat 7 7 is h8, got %s", res)
	}
}

func BenchmarkPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pos(uint8(i%8), uint8((i/8)%8))
	}
}

func BenchmarkPosConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		posConcat(uint8(i%8), uint8((i/8)%8))
	}
}

func pos(l, r uint8) string {
	return string([]byte{'a' + l, '1' + r})
}

func posConcat(l, r uint8) string {
	return string('a'+l) + string('1'+r)
}
