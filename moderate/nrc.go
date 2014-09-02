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
	for scanner.Scan() {
		s, seen := scanner.Text(), make(map[rune]int)
		for _, i := range s {
			seen[i]++
		}
		for _, i := range s {
			if seen[i] == 1 {
				fmt.Printf("%c", i)
				break
			}
		}
		fmt.Println()
	}
}
