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
		var (
			n uint64
			i int
		)
		fmt.Sscan(scanner.Text(), &n)
		for n > 0 {
			if n%2 > 0 {
				i++
			}
			n >>= 1
		}
		fmt.Println(i % 3)
	}
}
