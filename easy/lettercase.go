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
		var count, lower, ratio float64
		for _, i := range scanner.Text() {
			if i >= 'a' && i <= 'z' {
				lower++
				count++
			} else if i >= 'A' && i <= 'Z' {
				count++
			}
		}
		if count > 0 {
			ratio = 100 * lower / count
			fmt.Printf("lowercase: %.2f uppercase: %.2f\n", ratio, 100-ratio)
		} else {
			fmt.Println("lowercase: 0.00 uppercase: 0.00")
		}
	}
}
