package main

import (
	"fmt"
	"testing"
)

type tuple struct {
	r, g, b int
}

func TestColorCode(t *testing.T) {
	for k, v := range map[string]tuple{
		"(0.56,0.94,0.21,0.02)": tuple{110, 15, 197},
		"HSL(359,0,0)":          tuple{0, 0, 0},
		"HSV(276,33,7)":         tuple{15, 12, 18},
		"#cfa9c4":               tuple{207, 169, 196},
		"HSL(98,100,50)":        tuple{94, 255, 0},
		"#5eff00":               tuple{94, 255, 0}} {
		if r, g, b := colorCode(k); r != v.r || g != v.g || b != v.b {
			t.Errorf("failed: colorCode %s is %d %d %d, got %d %d %d",
				k, v.r, v.g, v.b, r, g, b)
		}
	}
}

func abs(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func mod2(a float32) float32 {
	return a - float32(int(a/2)*2)
}

func rint(a float32) int {
	return int(a*255 + 0.5)
}

func colorCode(q string) (r, g, b int) {
	switch {
	case q[0] == '#':
		fmt.Sscanf(q, "#%2x%2x%2x", &r, &g, &b)
	case q[:3] == "HSL":
		var h, s, l float32
		fmt.Sscanf(q, "HSL(%f,%f,%f)", &h, &s, &l)
		s, l, h = s/100, l/100, h/60
		c := (1 - abs(2*l-1)) * s
		x, m := c*(1-abs(float32(mod2(h)-1))), l-c/2
		switch {
		case h < 1:
			r, g, b = rint(c+m), rint(x+m), rint(m)
		case h < 2:
			r, g, b = rint(x+m), rint(c+m), rint(m)
		case h < 3:
			r, g, b = rint(m), rint(c+m), rint(x+m)
		case h < 4:
			r, g, b = rint(m), rint(x+m), rint(c+m)
		case h < 5:
			r, g, b = rint(x+m), rint(m), rint(c+m)
		case h < 6:
			r, g, b = rint(c+m), rint(m), rint(x+m)
		}
	case q[:3] == "HSV":
		var h, s, v float32
		fmt.Sscanf(q, "HSV(%f,%f,%f)", &h, &s, &v)
		s, v, h = s/100, v/100, h/60
		c := v * s
		x, m := c*(1-abs(float32(mod2(h)-1))), v-c
		switch {
		case h < 1:
			r, g, b = rint(c+m), rint(x+m), rint(m)
		case h < 2:
			r, g, b = rint(x+m), rint(c+m), rint(m)
		case h < 3:
			r, g, b = rint(m), rint(c+m), rint(x+m)
		case h < 4:
			r, g, b = rint(m), rint(x+m), rint(c+m)
		case h < 5:
			r, g, b = rint(x+m), rint(m), rint(c+m)
		case h < 6:
			r, g, b = rint(c+m), rint(m), rint(x+m)
		}
	case q[0] == '(':
		var c, m, y, k float32
		fmt.Sscanf(q, "(%f,%f,%f,%f)", &c, &m, &y, &k)
		r, g, b = rint((1-c)*(1-k)), rint((1-m)*(1-k)), rint((1-y)*(1-k))
	}
	return r, g, b
}
