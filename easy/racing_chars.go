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

	scanner.Scan()
	s := scanner.Text()
	curr := strings.Index(s, "C")
	if curr == -1 {
		curr = strings.Index(s, "_")
	}
	fmt.Println(s[:curr] + "|" + s[curr+1:])
	var px int
	for scanner.Scan() {
		px = curr
		s = scanner.Text()
		if len(s) < 1 {
			continue
		}
		a, b := px-1, px+2
		if a < 0 {
			a = 0
		}
		if b > len(s) {
			b = len(s)
		}
		curr = strings.Index(s[a:b], "C")
		if curr == -1 {
			curr = a + strings.Index(s[a:b], "_")
		} else {
			curr += a
		}
		if curr < px {
			s = s[:curr] + "/" + s[curr+1:]
		} else if curr == px {
			s = s[:curr] + "|" + s[curr+1:]
		} else {
			s = s[:curr] + `\` + s[curr+1:]
		}
		fmt.Println(s)
	}
}
