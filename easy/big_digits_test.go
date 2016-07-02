package main

import (
	"fmt"
	"strings"
	"testing"
)

const (
	digits = `
-**----*--***--***---*---****--**--****--**---**--
*--*--**-----*----*-*--*-*----*-------*-*--*-*--*-
*--*---*---**---**--****-***--***----*---**---***-
*--*---*--*-------*----*----*-*--*--*---*--*----*-
-**---***-****-***-----*-***---**---*----**---**--
--------------------------------------------------
`
	w, h = 5, 6
)

func TestBigDigits(t *testing.T) {
	for k, v := range map[string]string{
		"0": `-**--
*--*-
*--*-
*--*-
-**--
-----`,
		"9": `-**--
*--*-
-***-
---*-
-**--
-----`,
		"3.1415926": `***----*---*-----*--****--**--***---**--
---*--**--*--*--**--*----*--*----*-*----
-**----*--****---*--***---***--**--***--
---*---*-----*---*-----*----*-*----*--*-
***---***----*--***-***---**--****--**--
----------------------------------------`,
		"1.41421356": `--*---*-----*---*---***----*--***--****--**--
-**--*--*--**--*--*----*--**-----*-*----*----
--*--****---*--****--**----*---**--***--***--
--*-----*---*-----*-*------*-----*----*-*--*-
-***----*--***----*-****--***-***--***---**--
---------------------------------------------`,
		"01-01-1970": `-**----*---**----*----*---**--****--**--
*--*--**--*--*--**---**--*--*----*-*--*-
*--*---*--*--*---*----*---***---*--*--*-
*--*---*--*--*---*----*-----*--*---*--*-
-**---***--**---***--***--**---*----**--
----------------------------------------`,
		"2.7182818284": `***--****---*---**--***---**----*---**--***---**---*---
---*----*--**--*--*----*-*--*--**--*--*----*-*--*-*--*-
-**----*----*---**---**---**----*---**---**---**--****-
*-----*-----*--*--*-*----*--*---*--*--*-*----*--*----*-
****--*----***--**--****--**---***--**--****--**-----*-
-------------------------------------------------------`,
		"4 8 15 16 23 42": `-*----**----*--****---*---**--***--***---*---***--
*--*-*--*--**--*-----**--*-------*----*-*--*----*-
****--**----*--***----*--***---**---**--****--**--
---*-*--*---*-----*---*--*--*-*-------*----*-*----
---*--**---***-***---***--**--****-***-----*-****-
--------------------------------------------------`,
		"abcdefghijklmnop": `




`,
		"1234567890123456": `--*--***--***---*---****--**--****--**---**---**----*--***--***---*---****--**--
-**-----*----*-*--*-*----*-------*-*--*-*--*-*--*--**-----*----*-*--*-*----*----
--*---**---**--****-***--***----*---**---***-*--*---*---**---**--****-***--***--
--*--*-------*----*----*-*--*--*---*--*----*-*--*---*--*-------*----*----*-*--*-
-***-****-***-----*-***---**---*----**---**---**---***-****-***-----*-***---**--
--------------------------------------------------------------------------------`} {
		if r := bigDigits(k); r != v {
			t.Errorf("failed: bigDigits %s is\n%s\n got\n%s",
				k, v, r)
		}
	}
}

func BenchmarkBigDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigDigits(fmt.Sprint(i))
	}
}

func bigDigits(q string) string {
	r := make([]string, h)
	for _, i := range q {
		if i >= '0' && i <= '9' {
			for j := 0; j < h; j++ {
				pos := 1 + j*w*10 + j + int(i-'0')*w
				r[j] += digits[pos : pos+w]
			}
		}
	}
	return strings.Join(r, "\n")
}
