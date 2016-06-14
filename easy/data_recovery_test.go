package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRecovery(t *testing.T) {
	h := map[string]string{
		"2000 and was not However, implemented 1998 it until;9 8 3 4 1 5 7 2":                   "However, it was not implemented until 1998 and 2000",
		"programming first The language;3 2 1":                                                  "The first programming language",
		"programs Manchester The written ran Mark 1952 1 in Autocode from;6 2 1 7 5 3 11 4 8 9": "The Manchester Mark 1 ran programs written in Autocode from 1952"}
	for k, v := range h {
		s := strings.Split(k, ";")
		if res := recovery(s[0], s[1]); res != v {
			t.Errorf("failed: recovery\n %s\nis\n %s\ngot\n %s", k, v, res)
		}
	}
}

func recovery(s, t string) string {
	var k int
	w := strings.Fields(s)
	r := make([]string, len(w))
	for ix, i := range strings.Fields(t) {
		fmt.Sscan(i, &k)
		r[k-1] = w[ix]
	}
	for ix := range r {
		if r[ix] == "" {
			r[ix] = w[len(w)-1]
			break
		}
	}
	return strings.Join(r, " ")
}
