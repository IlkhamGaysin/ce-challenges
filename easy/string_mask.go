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
		s := strings.Fields(scanner.Text())
		for ix, i := range s[0] {
			if s[1][ix] == '1' {
				fmt.Print(strings.ToUpper(string(i)))
			} else {
				fmt.Printf("%c", i)
			}
		}
		fmt.Println()
	}
}
