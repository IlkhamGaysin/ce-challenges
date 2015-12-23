package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		u bool
		c byte = '\n'
	)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		c = scanner.Text()[0]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			u = !u
			if u {
				c &= 223
			} else {
				c |= 32
			}
		} else {
			if c == '\n' {
				u = false
			}
		}
		fmt.Print(string(c))
	}
	if c != '\n' {
		fmt.Println()
	}
}
