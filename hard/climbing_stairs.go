package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var n int
		fmt.Sscan(scanner.Text(), &n)
		if n < 2 {
			fmt.Println(n)
		} else {
			r := big.NewInt(1)
			q := big.NewInt(0)
			for p := big.NewInt(1); n > 1; n-- {
				q.Set(r)
				r.Add(r, p)
				p.Set(q)
			}
			fmt.Println(r)
		}
	}
}
