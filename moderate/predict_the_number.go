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
		var n, i uint
		fmt.Sscan(scanner.Text(), &n)
		for ; n > 0; n &= n - 1 {
			i++
		}
		fmt.Println(i % 3)
	}
}
