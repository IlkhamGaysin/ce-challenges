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
		maph := func(r rune) rune {
			if strings.Contains("0123456789", string(r)) {
				return r
			} else if strings.Contains("abcdefghij", string(r)) {
				return r - 'a' + '0'
			}
			return -1
		}
		t := strings.Map(maph, scanner.Text())

		if len(t) == 0 {
			fmt.Println("NONE")
		} else {
			fmt.Println(t)
		}
	}
}
