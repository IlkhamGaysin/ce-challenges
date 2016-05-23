package main

import "testing"

func TestMask(t *testing.T) {
	if res := mask("hello", "11001"); res != "HEllO" {
		t.Errorf("failed: mask hello 11001 is HEllO, got %s", res)
	}
	if res := mask("world", "10000"); res != "World" {
		t.Errorf("failed: mask world 10000 is World, got %s", res)
	}
	if res := mask("cba", "111"); res != "CBA" {
		t.Errorf("failed: mask cba 111 is CBA, got %s", res)
	}
}

func BenchmarkMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, t []byte
		for j := i; j > 0; j /= 26 {
			s = append(s, byte(j%26+97))
			t = append(t, '1')
		}
		mask(string(s), string(t))
	}
}

func upper(b byte) byte {
	return b & 223
}

func mask(s, t string) string {
	r := make([]byte, len(s))
	for ix, i := range s {
		if t[ix] == '1' {
			r[ix] = upper(byte(i))
		} else {
			r[ix] = byte(i)
		}
	}
	return string(r)
}
