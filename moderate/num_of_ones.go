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
		var a, c int
		fmt.Sscan(scanner.Text(), &a)
		for ; a > 0; a &= a - 1 {
			c++
		}
		fmt.Println(c)
	}
}
