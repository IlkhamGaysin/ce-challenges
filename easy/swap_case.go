package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var c byte = '\n'
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		c = scanner.Text()[0]
		switch {
		case c >= 'a' && c <= 'z':
			c &= 223
		case c >= 'A' && c <= 'Z':
			c |= 32
		}
		fmt.Printf("%c", c)
	}
	if c != '\n' {
		fmt.Println()
	}
}
