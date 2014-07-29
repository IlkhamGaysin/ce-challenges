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
	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		var count, lower, ratio float64
		for _, i := range s {
			if i >= 'a' && i <= 'z' {
				lower++
				count++
			} else if i >= 'A' && i <= 'Z' {
				count++
			}
		}
		ratio = 100 * lower / count
		fmt.Printf("lowercase: %.2f uppercase: %.2f\n", ratio, 100-ratio)
	}
}
