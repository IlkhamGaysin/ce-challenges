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
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var n, r int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		p := make([]int, len(scanner.Text()))
		for m := n; m > 0; m /= 10 {
			if v := m % 10; v < len(p) {
				p[v]++
			} else {
				goto out
			}
		}
		for _, i := range p {
			r = r*10 + i
		}
		if n == r {
			fmt.Println(1)
			continue
		}
	out:
		fmt.Println(0)
	}
}
