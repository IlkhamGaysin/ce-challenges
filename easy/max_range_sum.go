package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		t := strings.Fields(s[1])
		var m, n, c int
		fmt.Sscan(s[0], &n)
		u := make([]int, len(t))
		for ix, i := range t {
			fmt.Sscan(i, &u[ix])
		}
		for i := 0; i < n; i++ {
			c += u[i]
		}
		if c > m {
			m = c
		}
		for len(u) > n {
			c = c - u[0] + u[n]
			if c > m {
				m = c
			}
			u = u[1:]
		}
		fmt.Println(m)
	}
}
