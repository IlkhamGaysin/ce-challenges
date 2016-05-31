package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMultiples(t *testing.T) {
	h := map[string]string{
		"1,2,3,4;4,5,6":               "4",
		"7,8,9;8,9,10,11,12":          "8,9",
		"3,4,7;2,6,12":                "",
		"1,3,5,7,9,11,13;2,3,5,6,8,9": "3,5,9"}
	for k, v := range h {
		s := strings.Split(k, ";")
		if res := intersect(s[0], s[1]); res != v {
			t.Errorf("failed: intersect %s is %s, got %s", k, v, res)
		}
	}
}

func intersect(s, t string) string {
	var (
		a    int
		x, y []int
		r    []string
	)
	u, v := strings.Split(s, ","), strings.Split(t, ",")
	for _, i := range u {
		fmt.Sscan(i, &a)
		x = append(x, a)
	}
	for _, i := range v {
		fmt.Sscan(i, &a)
		y = append(y, a)
	}
	for i, j := 0, 0; i < len(u) && j < len(v); {
		if x[i] == y[j] {
			r = append(r, fmt.Sprint(x[i]))
			i++
			j++
		} else if x[i] > y[j] {
			j++
		} else {
			i++
		}
	}
	return strings.Join(r, ",")
}
