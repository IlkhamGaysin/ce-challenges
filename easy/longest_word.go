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
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Split(strings.TrimSpace(s), " ")
		maxw, maxl := "", 0
		for _, i := range t {
			if len(i) > maxl {
				maxw, maxl = i, len(i)
			}
		}
		fmt.Println(maxw)
	}
}
