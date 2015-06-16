package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func contains(a, b string) (string, bool) {
	res := strings.Repeat("_", len(a))
	i := strings.Index(a, b)
	if i == -1 {
		return res, false
	}
	return res[:i] + b + res[i+len(b):], true
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var (
			c   int
			res []string
		)
		s := strings.Split(scanner.Text(), ";")
		t, u := strings.Fields(s[0]), strings.Fields(s[1])
		for i := 0; i < len(t); i++ {
			var r string
			if c < len(u) {
				var f bool
				if r, f = contains(t[i], u[c]); f {
					c++
				}
			} else {
				r = strings.Repeat("_", len(t[i]))
			}
			if len(r) > 0 {
				res = append(res, r)
			}
		}
		if c < len(u) {
			fmt.Println("I cannot fix history")
		} else {
			fmt.Println(strings.Join(res, " "))
		}
	}
}
