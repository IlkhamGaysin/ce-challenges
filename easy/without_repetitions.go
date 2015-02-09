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
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	var c string
	for scanner.Scan() {
		if c != scanner.Text() {
			fmt.Print(scanner.Text())
		}
		c = scanner.Text()
	}
	if c != "\n" {
		fmt.Println()
	}
}
