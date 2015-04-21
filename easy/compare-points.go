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
		var x1, y1, x2, y2 int
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &x1, &y1, &x2, &y2)
		switch {
		case x1 == x2:
			if y1 == y2 {
				fmt.Println("here")
			} else if y1 < y2 {
				fmt.Println("N")
			} else {
				fmt.Println("S")
			}
		case y1 == y2:
			if x1 < x2 {
				fmt.Println("E")
			} else {
				fmt.Println("W")
			}
		case x1 < x2:
			if y1 < y2 {
				fmt.Println("NE")
			} else {
				fmt.Println("SE")
			}
		default:
			if y1 < y2 {
				fmt.Println("NW")
			} else {
				fmt.Println("SW")
			}
		}
	}
}
