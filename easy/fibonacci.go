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
		if n < 2 {
			fmt.Println(n)
		} else {
			r := 1
			for p := 1; n > 2; n-- {
				r, p = r+p, r
			}
			fmt.Println(r)
		}
	}
}
