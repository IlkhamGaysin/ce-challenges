package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var u bool
	c := '\n'
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		c = rune(scanner.Text()[0])
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			u = !u
			if u {
				fmt.Print(strings.ToUpper(string(c)))
			} else {
				fmt.Print(strings.ToLower(string(c)))
			}
		} else {
			if c == '\n' {
				u = false
			}
			fmt.Print(string(c))
		}
	}
	if c != '\n' {
		fmt.Println()
	}
}
