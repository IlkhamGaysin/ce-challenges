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
		var n, p, q uint
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &n, &p, &q)
		fmt.Println(((n & (1 << (p - 1))) == 0) == ((n & (1 << (q - 1))) == 0))
	}
}
