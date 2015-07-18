package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isAlpha(c string) bool {
	if (c[0] >= 'a' && c[0] <= 'z') || (c[0] >= 'A' && c[0] <= 'Z') {
		return true
	}
	return false
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		c, w := string(scanner.Text()[0]), true
		for !isAlpha(c) {
			if scanner.Scan() {
				c = string(scanner.Text()[0])
			} else {
				c = "\n"
				break
			}
		}
		for c != "\n" {
			if isAlpha(c) {
				if !w {
					fmt.Printf(" ")
					w = true
				}
				fmt.Printf(strings.ToLower(c))
			} else if w {
				w = false
			}
			if scanner.Scan() {
				c = string(scanner.Text()[0])
			} else {
				c = "\n"
			}
		}
		fmt.Println()
		w = true
	}
}
