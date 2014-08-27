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
		var c string
		t := []string{}
		for _, i := range s {
			if i == ' ' {
				t = append([]string{c}, t...)
				c = ""
			} else {
				c += string(i)
			}
		}
		t = append([]string{c}, t...)
		fmt.Println(strings.Join(t, " "))
	}
}
