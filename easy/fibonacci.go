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
		var n int
		fmt.Sscanf(string(s), "%d", &n)
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
