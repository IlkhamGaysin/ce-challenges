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
		var c int
		for i, s := 0, scanner.Text(); i < len(s)-4; i++ {
			if s[i:i+5] == ">>-->" || s[i:i+5] == "<--<<" {
				c++
			}
		}
		fmt.Println(c)
	}
}
