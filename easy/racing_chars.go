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

	s, _ := reader.ReadString('\n')
	curr := strings.Index(s, "C")
	if curr == -1 {
		curr = strings.Index(s, "_")
	}
	fmt.Print(s[:curr] + "|" + s[curr+1:])
	var px int
	for {
		px = curr
		s, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		a, b := px-1, px+2
		if a < 0 {
			a = 0
		}
		if b > len(s)-1 {
			b = len(s) - 1
		}
		curr = strings.Index(s[a:b], "C")
		if curr == -1 {
			curr = px + strings.Index(s[a:b], "_") - 1
		} else {
			curr += px - 1
		}
		if curr < px {
			s = s[:curr] + "/" + s[curr+1:]
		} else if curr == px {
			s = s[:curr] + "|" + s[curr+1:]
		} else {
			s = s[:curr] + "\\" + s[curr+1:]
		}
		fmt.Print(s)
	}
}
