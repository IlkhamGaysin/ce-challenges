package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	r := map[int]bool{
		0:        true,
		2:        true,
		10:       false,
		11:       true,
		347:      false,
		12621:    true,
		3490943:  true,
		12345670: false}
	for k, v := range r {
		if res := isPalindrome(k); res != v {
			t.Errorf("failed: isPalindrome %d is %t, got %t", k, v, res)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(i)
	}
}

func isPalindrome(a int) bool {
	if a%10 == 0 {
		if a > 0 {
			return false
		}
		return true
	}
	var rev int
	for ; a > rev; a /= 10 {
		rev = 10*rev + a%10
		if a == rev {
			return true
		}
	}
	return rev == a
}
