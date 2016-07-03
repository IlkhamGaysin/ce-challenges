package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var sum, v int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &v)
		sum += v
	}
	fmt.Println(sum)
}
