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
	var col int
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &col)
		r := ""
		for col > 0 {
			col--
			r = string('A'+col%26) + r
			col /= 26
		}
		fmt.Println(r)
	}
}
