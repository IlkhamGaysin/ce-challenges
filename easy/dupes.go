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
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), ",")
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
