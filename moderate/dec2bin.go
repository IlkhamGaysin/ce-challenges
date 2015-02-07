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
		if len(scanner.Text()) > 0 {
			var t int
			fmt.Sscanf(scanner.Text(), "%d", &t)
			fmt.Printf("%b\n", t)
		}
	}
}
