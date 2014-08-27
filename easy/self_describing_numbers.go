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
		var n, r int
		fmt.Sscanf(string(s), "%d", &n)
		p := make([]int, len(s))
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
