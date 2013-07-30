package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var primes = []int{2, 3, 5, 7, 11, 13}

func isPrime(a int) bool {
	for b := 0; a%primes[b] > 0; b++ {
		if primes[b]*primes[b] > a {
			primes = append(primes, a)
			return true
		}
	}
	return false
}

func primeSeq() func() int {
	p0 := 0
	var i int
	return func() int {
		if p0 < len(primes) {
			i = primes[p0]
			p0++
			return i
		}
		for i += 2; !isPrime(i); i += 2 {
		}
		p0++
		return i
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var a, b, count int
		fmt.Sscanf(s, "%d,%d", &a, &b)
		nextPrime := primeSeq()
		for i := nextPrime(); i <= b; i = nextPrime() {
			if i >= a {
				count++
			}
		}
		fmt.Println(count)
	}
}
