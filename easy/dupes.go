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
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		t := strings.Split(string(s), ",")
		var p string
		q := []string{}
		for _, i := range t {
			if i != p {
				q = append(q, i)
				p = i
			}
		}
		fmt.Println(strings.Join(q, ","))
	}
}
