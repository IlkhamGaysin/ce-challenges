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
	var u bool
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}
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
}
