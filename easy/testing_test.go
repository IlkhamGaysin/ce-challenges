package main

import (
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	for k, v := range map[string]string{
		"Heelo Codevval | Hello Codeeval": "Low",
		"hELLO cODEEVAL | Hello Codeeval": "Critical",
		"Hello Codeeval | Hello Codeeval": "Done"} {
		s := strings.Split(k, " | ")
		if r := test(s[0], s[1]); r != v {
			t.Errorf("failed: test %s is %s, got %s", k, v, r)
		}
	}
}

func test(s, t string) string {
	var r int
	for i := range s {
		if s[i] != t[i] {
			r++
			if r > 6 {
				break
			}
		}
	}
	switch {
	case r == 0:
		return "Done"
	case r <= 2:
		return "Low"
	case r <= 4:
		return "Medium"
	case r <= 6:
		return "High"
	default:
		return "Critical"
	}
}
