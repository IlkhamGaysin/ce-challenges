package main

import "testing"

func TestMultiples(t *testing.T) {
	if res := multiples(13, 8); res != 16 {
		t.Errorf("failed: multiples 13 8 is 16, got %d", res)
	}
	if res := multiples(17, 16); res != 32 {
		t.Errorf("failed: multiples 17 16 is 32, got %d", res)
	}
	if res := multiples(2000000000000000, 1); res != 2000000000000000 {
		t.Errorf("failed: multiples 2000000000000000 1 is 2000000000000000, got %d", res)
	}
}

func BenchmarkMultiples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiples(i/16+1, i%16+1)
	}
}

func BenchmarkMultiplesNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplesNaive(i/16+1, i%16+1)
	}
}

func multiples(x, n int) (r int) {
	var c uint
	r = n
	for (r << 1) <= x {
		r <<= 1
		c++
	}
	for c > 0 {
		c--
		for r+(n<<c) <= x {
			r += n << c
		}
	}
	for r < x {
		r += n
	}
	return r
}

func multiplesNaive(x, n int) (r int) {
	r = n
	for r < x {
		r += n
	}
	return r
}
