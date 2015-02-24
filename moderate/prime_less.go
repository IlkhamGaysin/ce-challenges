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
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var a int
		fmt.Sscan(scanner.Text(), &a)
		nextPrime := primeSeq()
		for i := nextPrime(); i < a; i = nextPrime() {
			if i > 2 {
				fmt.Print(",")
			}
			fmt.Print(i)
		}
		fmt.Println()
	}
}
