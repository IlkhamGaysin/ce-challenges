package main

import "testing"

func TestBlackCard(t *testing.T) {
	if res := blackCard([]string{"John", "Sara", "Tom", "Susan"}, 3); res != "Sara" {
		t.Errorf("failed: blackCard John Sara Tom Susan | 3 is Sara, got %s", res)
	}
	if res := blackCard([]string{"John", "Tom", "Mary"}, 5); res != "Mary" {
		t.Errorf("failed: blackCard John Tom Mary | 5 is Mary, got %s", res)
	}
}

func blackCard(s []string, m int) string {
	for len(s) > 1 {
		n := m%len(s) - 1
		if n == -1 {
			s = s[:len(s)-1]
		} else {
			s = append(s[:n], s[n+1:]...)
		}
	}
	return s[0]
}
