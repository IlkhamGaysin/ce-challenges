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
	reader := bufio.NewReader(data)
	var r int
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}
		if c >= '0' && c <= '9' {
			r += int(c - '0')
		} else {
			fmt.Println(r)
			r = 0
		}
	}
}
