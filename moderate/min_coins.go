package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	b := []int{0, 1, 2, 1, 2}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var a int
		fmt.Sscanf(scanner.Text(), "%d", &a)
		fmt.Println(a/5 + b[a%5])
	}
}
