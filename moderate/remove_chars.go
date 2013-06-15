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
		s = strings.TrimRight(s, ",")
		t, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		t = strings.TrimSpace(t)
		mapf := func(r rune) rune {
			if strings.Contains(t, string(r)) {
				return -1
			}
			return r
		}
		fmt.Println(strings.Map(mapf, s))
	}
}
