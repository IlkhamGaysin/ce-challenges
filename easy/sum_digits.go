package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var r int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		c := scanner.Text()[0]
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
