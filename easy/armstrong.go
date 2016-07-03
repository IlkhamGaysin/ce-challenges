package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func powi(y, x int) int {
	ret := 1
	for x > 0 {
		if x&1 > 0 {
			ret *= y
		}
		y *= y
		x >>= 1
	}
	return ret
}

func main() {
	var m int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var e, t int
		fmt.Sscan(scanner.Text(), &m)
		for a := m; a > 0; a /= 10 {
			e++
		}
		for a := m; a > 0; a /= 10 {
			t += powi(a%10, e)
		}
		if m == t {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
