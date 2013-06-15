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
	reader := bufio.NewReader(data)
	var s string
	for {
		s, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(strings.ToLower(s))
	}
}
