package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}
		switch {
		case c >= 'a' && c <= 'z':
			fmt.Print(strings.ToUpper(string(c)))
		case c >= 'A' && c <= 'Z':
			fmt.Print(strings.ToLower(string(c)))
		default:
			fmt.Printf("%c", c)
		}
	}
}
