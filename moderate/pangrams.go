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
		s := strings.ToLower(scanner.Text())
		var f bool
		for i := 'a'; i <= 'z'; i++ {
			if !strings.Contains(s, string(i)) {
				fmt.Print(string(i))
				f = true
			}
		}
		if !f {
			fmt.Println("NULL")
		} else {
			fmt.Println()
		}
	}
}
