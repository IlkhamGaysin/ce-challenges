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
		s := strings.Split(scanner.Text(), ";")
		m, b, been := map[string]string{}, "BEGIN", []string{}
		for _, i := range s {
			t := strings.Split(i, "-")
			if t[0] == t[1] {
				goto invalid
			}
			m[t[0]] = t[1]
		}
		for i := 0; i < len(s); i++ {
			var found bool
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
