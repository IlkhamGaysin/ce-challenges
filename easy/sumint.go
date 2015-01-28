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
	var sum int
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var v int
		fmt.Sscanf(scanner.Text(), "%d", &v)
		sum += v
	}
	fmt.Println(sum)
}
