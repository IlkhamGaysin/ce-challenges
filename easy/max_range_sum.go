package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func maxRangeSum(q string) (r int) {
	s := strings.Split(q, ";")
	t := strings.Fields(s[1])
	var n, c int
	fmt.Sscan(s[0], &n)
	u := make([]int, len(t))
	for ix, i := range t {
		fmt.Sscan(i, &u[ix])
	}
	for i := 0; i < n; i++ {
		c += u[i]
	}
	if c > r {
		r = c
	}
	for len(u) > n {
		c = c - u[0] + u[n]
		if c > r {
			r = c
		}
		u = u[1:]
	}
	return r
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Println(maxRangeSum(scanner.Text()))
	}
}
