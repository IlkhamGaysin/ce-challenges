package main

import (
	"fmt"
	"testing"
)

func TestMonth(t *testing.T) {
	r := map[string]int{
		"Feb 2004": 169,
		"Jul 2008": 222,
		"Aug 2011": 259,
		"Sep 2013": 284}
	for k, v := range r {
		if res := month(k); res != v {
			t.Errorf("failed: month %s is %d, got %d", k, v, res)
		}
	}
}

func BenchmarkMonth(b *testing.B) {
	mnth := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul",
		"Aug", "Sep", "Oct", "Nov", "Dec"}
	for i := 0; i < b.N; i++ {
		month(mnth[i%12] + " " + fmt.Sprint((i/12)%27+1990))
	}
}

var moy map[string]int

func init() {
	moy = map[string]int{"Jan": 0, "Feb": 1, "Mar": 2, "Apr": 3,
		"May": 4, "Jun": 5, "Jul": 6, "Aug": 7,
		"Sep": 8, "Oct": 9, "Nov": 10, "Dec": 11}
}

func month(s string) int {
	var k int
	fmt.Sscanf(s[4:8], "%d", &k)
	return 12*(k-1990) + moy[s[0:3]]
}
