package main

import "testing"

type tuple struct {
	l, r uint8
}

func TestPos(t *testing.T) {
	for k, v := range map[tuple]string{
		tuple{0, 0}: "a1",
		tuple{7, 7}: "h8"} {
		if r := pos(k.l, k.r); r != v {
			t.Errorf("failed: pos %d %d is %s, got %s", k.l, k.r, v, r)
		}
	}
}

func TestPosConcat(t *testing.T) {
	for k, v := range map[tuple]string{
		tuple{0, 0}: "a1",
		tuple{7, 7}: "h8"} {
		if r := posConcat(k.l, k.r); r != v {
			t.Errorf("failed: posConcat %d %d is %s, got %s", k.l, k.r, v, r)
		}
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
