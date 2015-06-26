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
		var (
			maxw string
			maxl int
		)
		t := strings.Fields(strings.TrimSpace(scanner.Text()))
		for _, i := range t {
			if len(i) > maxl {
				maxw, maxl = i, len(i)
			}
		}
		r := make([]string, maxl)
		for i := 0; i < maxl; i++ {
			r[i] = strings.Repeat("*", i) + string(maxw[i])
		}
		fmt.Println(strings.Join(r, " "))
	}
}
