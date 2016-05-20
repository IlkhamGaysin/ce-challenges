package main

import (
	"strings"
	"testing"
)

func TestPangram(t *testing.T) {
	h := map[string]string{
		"A quick brown fox jumps over the lazy dog":        "NULL",
		"A slow Yellow fox crawls Under the proactive dog": "bjkmqz"}
	for k, v := range h {
		if res := pangram(k); res != v {
			t.Errorf("failed: pangram\n %s\nis\n %s,\ngot\n %s", k, v, res)
		}
	}
}

func BenchmarkPangram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []byte
		for j := i; j > 0; j /= 26 {
			s = append(s, byte(j%26+97))
		}
		pangram(string(s))
	}
}

func pangram(s string) (r string) {
	s = strings.ToLower(s)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(s, string(i)) {
			r += string(i)
		}
	}
	if r == "" {
		r = "NULL"
	}
	return r
}
