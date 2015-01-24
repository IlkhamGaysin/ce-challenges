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
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Fields(strings.TrimSpace(scanner.Text()))
		maxw, maxl := "", 0
		for _, i := range t {
			if len(i) > maxl {
				maxw, maxl = i, len(i)
			}
		}
		fmt.Println(maxw)
	}
}
