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
		if len(scanner.Text()) > 0 {
			var (
				c string
				t []string
			)
			for _, i := range scanner.Text() {
				if i == ' ' {
					t, c = append([]string{c}, t...), ""
				} else {
					c += string(i)
				}
			}
			t = append([]string{c}, t...)
			fmt.Println(strings.Join(t, " "))
		}
	}
}
