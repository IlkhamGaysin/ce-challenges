package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var a, b, x, y int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var r int
		fmt.Sscanf(scanner.Text(), "%dx%d | %d %d", &a, &b, &x, &y)
		for b != y {
			r, a, b, x, y = r+a, b-1, a, b-y, x
		}
		fmt.Println(r + x)
	}
}
