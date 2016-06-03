package main

import (
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	h := map[string]string{
		"Heelo Codevval | Hello Codeeval": "Low",
		"hELLO cODEEVAL | Hello Codeeval": "Critical",
		"Hello Codeeval | Hello Codeeval": "Done"}
	for k, v := range h {
		s := strings.Split(k, " | ")
		if res := test(s[0], s[1]); res != v {
			t.Errorf("failed: test %s is %s, got %s", k, v, res)
		}
	}
}

func test(s, t string) string {
	var r int
	for i := 0; i < len(s); i++ {
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
