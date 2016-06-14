package main

import (
	"strings"
	"testing"
)

func TestLongest(t *testing.T) {
	h := map[string]string{
		"some line with text": "some",
		"another line":        "another"}
	for k, v := range h {
		if r := longest(k); r != v {
			t.Errorf("failed: longest %s is %s, got %s", k, v, r)
		}
	}
}

func longest(q string) (r string) {
	for _, i := range strings.Fields(q) {
		if len(i) > len(r) {
			r = i
		}
	}
	return r
}
