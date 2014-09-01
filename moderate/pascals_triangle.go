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
		var n int
		fmt.Sscanf(scanner.Text(), "%d", &n)
		fmt.Print("1")
		for i := 1; i < n; i++ {
			fmt.Print(" 1")
			r := 1
			for j := 1; j <= i; j++ {
				r = (r * (i + 1 - j)) / j
				fmt.Printf(" %d", r)
			}
		}
		fmt.Println()
	}
}
