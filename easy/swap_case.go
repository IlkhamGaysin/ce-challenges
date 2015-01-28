package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
		switch {
		case c >= 'a' && c <= 'z':
			fmt.Print(strings.ToUpper(scanner.Text()))
		case c >= 'A' && c <= 'Z':
			fmt.Print(strings.ToLower(scanner.Text()))
		default:
			fmt.Printf("%c", c)
		}
	}
	if c != '\n' {
		fmt.Println()
	}
}
