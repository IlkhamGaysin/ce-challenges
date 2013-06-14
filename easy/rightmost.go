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
	var s, t string
	for {
		s, err = reader.ReadString(',')
		if err != nil {
			break
		}
		t, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.LastIndex(s, strings.TrimSpace(t)))
	}
}
