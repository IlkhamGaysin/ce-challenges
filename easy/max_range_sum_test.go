package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMaxRangeSum(t *testing.T) {
	h := map[string]int{
		"5;7 -3 -10 4 2 8 -2 4 -5 -2": 16,
		"6;-4 3 -10 5 3 -7 -3 7 -6 3": 0,
		"3;-7 0 -45 34 -24 7":         17}
	for k, v := range h {
		if r := maxRangeSum(k); r != v {
			t.Errorf("failed: maxRangeSum %s is %d, got %d", k, v, r)
		}
	}
}

func maxRangeSum(q string) (r int) {
	s := strings.Split(q, ";")
	t := strings.Fields(s[1])
	var n, c int
	fmt.Sscan(s[0], &n)
	u := make([]int, len(t))
	for ix, i := range t {
		fmt.Sscan(i, &u[ix])
	}
	for i := 0; i < n; i++ {
		c += u[i]
	}
	if c > r {
		r = c
	}
	for len(u) > n {
		c = c - u[0] + u[n]
		if c > r {
			r = c
		}
		u = u[1:]
	}
	return r
}
