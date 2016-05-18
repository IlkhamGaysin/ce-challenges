package main

import (
	"strings"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	h := map[string]bool{":))": false, "(:)": true, "j": true, "": true, ":):": true,
		"( :) ))": false, "asdf(asdf:)ghjk:ghjk(ghjk)": true, "(((::))()": false}
	for k, v := range h {
		if res := isBalanced(k, 0); res != v {
			t.Errorf("failed: isBalanced %s is %t, got %t", k, v, res)
		}
	}
}

func TestIsBalanced2(t *testing.T) {
	h := map[string]bool{":))": false, "(:)": true, "j": true, "": true, ":):": true,
		"( :) ))": false, "asdf(asdf:)ghjk:ghjk(ghjk)": true, "(((::))()": false}
	for k, v := range h {
		if res := isBalanced2(strings.TrimRight(k, "abcdefghijklmnopqrstuvwxyz :"), 0); res != v {
			t.Errorf("failed: isBalanced %s is %t, got %t", k, v, res)
		}
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	ch := []string{"(", ")", ":", "j", " "}
	for i := 0; i < b.N; i++ {
		var (
			s string
			j int
		)
		for j > 0 {
			s += ch[j%5]
			j /= 5
		}
		isBalanced(s, 0)
	}
}

func BenchmarkIsBalanced2(b *testing.B) {
	ch := []string{"(", ")", ":", "j", " "}
	for i := 0; i < b.N; i++ {
		var (
			s string
			j int
		)
		for j > 0 {
			s += ch[j%5]
			j /= 5
		}
		isBalanced2(strings.TrimRight(s, "abcdefghijklmnopqrstuvwxyz :"), 0)
	}
}

func isBalanced(s string, c int) bool {
	for s != "" && c >= 0 {
		first, last := s[0], s[len(s)-1]
		switch {
		case (first >= 'a' && first <= 'z') || first == ' ':
			s = s[1:]
		case (last >= 'a' && last <= 'z') || last == ' ' || last == ':':
			s = s[:len(s)-1]
		case first == '(':
			c, s = c+1, s[1:]
		case first == ')':
			c, s = c-1, s[1:]
		case first == ':':
			if s[1] == '(' {
				return isBalanced(s[2:], c) || isBalanced(s[2:], c+1)
			} else if s[1] == ')' {
				return isBalanced(s[2:], c) || isBalanced(s[2:], c-1)
			} else {
				s = s[1:]
			}
		default:
			return false
		}
	}
	return c == 0
}

func isBalanced2(s string, c int) bool {
	for s != "" && c >= 0 {
		first := s[0]
		switch {
		case (first >= 'a' && first <= 'z') || first == ' ':
			s = s[1:]
		case first == '(':
			c, s = c+1, s[1:]
		case first == ')':
			c, s = c-1, s[1:]
		case first == ':':
			if s[1] == '(' {
				return isBalanced2(s[2:], c) || isBalanced2(s[2:], c+1)
			} else if s[1] == ')' {
				return isBalanced2(s[2:], c) || isBalanced2(s[2:], c-1)
			} else {
				s = s[1:]
			}
		default:
			return false
		}
	}
	return c == 0
}
