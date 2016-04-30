package main

import (
	"strings"
	"testing"
)

func TestContainsAll(t *testing.T) {
	if !containsAll("ddccbbaa", "abcd") {
		t.Error("failed: ddccbbaa, abcd")
	}
	if !containsAll("Chardonnay", "ann") {
		t.Error("failed: Chardonnay, ann")
	}
	if containsAll("qwer", "ee") {
		t.Error("failed: qwer, ee")
	}
	if containsAll("Cabernet", "ot") {
		t.Error("failed: Cabernet, ot")
	}
}

func BenchmarkContainsAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsAll("Chardonnay", "ann")
		containsAll("Cabernet", "ot")
	}
}

func containsAll(a, b string) bool {
	for _, i := range b {
		ix := strings.IndexRune(a, i)
		if ix == -1 {
			return false
		}
		a = a[:ix] + a[ix+1:]
	}
	return true
}
