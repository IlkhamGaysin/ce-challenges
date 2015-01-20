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
		var m int
		fmt.Sscanf(scanner.Text(), "%d", &m)
		if m%2 != 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
