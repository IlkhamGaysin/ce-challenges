package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var found bool
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var been []string
		s := strings.Split(scanner.Text(), ";")
		m, b := make(map[string]string, len(s)), "BEGIN"
		for _, i := range s {
			t := strings.Split(i, "-")
			if t[0] == t[1] {
				goto invalid
			}
			m[t[0]] = t[1]
		}
		for _ = range s {
			if b, found = m[b]; !found {
				goto invalid
			}
			for _, j := range been {
				if b == j {
					goto invalid
				}
			}
			been = append(been, b)
		}
		if b == "END" {
			fmt.Println("GOOD")
			continue
		}
	invalid:
		fmt.Println("BAD")
	}
}
