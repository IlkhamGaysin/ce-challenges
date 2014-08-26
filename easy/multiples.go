package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		var x, n, r int
		var c uint
		fmt.Sscanf(string(s), "%d,%d", &x, &n)
		r = n
		for (r << 1) <= x {
			r <<= 1
			c++
		}
		for c > 0 {
			c--
			for r + (n << c) <= x {
				r += n << c
			}
		}
		for r < x {
			r += n
		}
		fmt.Println(r)
	}
}

