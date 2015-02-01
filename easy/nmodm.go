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
		var n, m int
		fmt.Sscanf(scanner.Text(), "%d,%d", &n, &m)
		fmt.Println(n - (n/m)*m)
	}
}
