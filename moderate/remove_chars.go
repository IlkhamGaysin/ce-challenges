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
		s := strings.Split(scanner.Text(), ", ")
		mapf := func(r rune) rune {
			if strings.Contains(s[1], string(r)) {
				return -1
			}
			return r
		}
		fmt.Println(strings.Map(mapf, s[0]))
	}
}
