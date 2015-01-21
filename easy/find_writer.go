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
		if scanner.Text() == "" {
			continue
		}
		t := strings.Split(scanner.Text(), "|")
		u := strings.Fields(strings.TrimSpace(t[1]))
		for _, i := range u {
			var f int
			fmt.Sscan(i, &f)
			fmt.Printf("%c", t[0][f-1])
		}
		fmt.Println()
	}
}
