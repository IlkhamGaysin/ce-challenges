package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var m int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &m)
		fmt.Println(1 - m&1)
	}
}
