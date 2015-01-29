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
	scanner.Split(bufio.ScanBytes)
	var r int
	for scanner.Scan() {
		c := rune(scanner.Text()[0])
		if c >= '0' && c <= '9' {
			r += int(c - '0')
		} else {
			fmt.Println(r)
			r = 0
		}
	}
	if r > 0 {
		fmt.Println(r)
	}
}
