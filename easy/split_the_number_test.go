package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMask(t *testing.T) {
	if res := split("3413289830", "a-bcdefghij"); res != -413289827 {
		t.Errorf("failed: split 3413289830 a-bcdefghij is -413289827, got %d", res)
	}
	if res := split("776", "a+bc"); res != 83 {
		t.Errorf("failed: split 776 a+bc is 83, got %d", res)
	}
	if res := split("12345", "a+bcde"); res != 2346 {
		t.Errorf("failed: split 12345 a+bcde is 2346, got %d", res)
	}
	if res := split("1232", "ab+cd"); res != 44 {
		t.Errorf("failed: split 1232 ab+cd is 44, got %d", res)
	}
	if res := split("90602", "a+bcde"); res != 611 {
		t.Errorf("failed: split 90602 a+bcde is 611, got %d", res)
	}
}

func split(s, t string) int {
	var n, m, p int
	var neg bool
	if strings.Contains(t, "+") {
		p = strings.Index(t, "+")
	} else {
		p = strings.Index(t, "-")
		neg = true
	}
	fmt.Sscanf(s[0:p], "%d", &n)
	fmt.Sscanf(s[p:len(s)], "%d", &m)
	if neg {
		return n - m
	}
	return n + m
}
