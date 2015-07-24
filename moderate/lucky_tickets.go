package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	var m, i int64
	r, b, c := new(big.Int), new(big.Int), new(big.Int)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &m)
		a := true

		r.SetInt64(0)
		for i = 0; i < m/2; i++ {
			b.Binomial(m, i)
			c.Binomial(11*m/2-1-10*i, m-1)
			if a {
				r.Add(r, b.Mul(b, c))
			} else {
				r.Sub(r, b.Mul(b, c))
			}
			a = !a
		}

		fmt.Println(r)
	}
}
