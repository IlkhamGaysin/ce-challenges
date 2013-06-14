package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		m, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatal(err)
		}
		if m%2 != 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
