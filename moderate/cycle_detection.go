package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func uniq(s []string, t string) bool {
	for _, i := range s {
		if i == t {
			return false
		}
	}
	return true
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		for len(s) > 1 && uniq(s[1:len(s)], s[0]) {
			s = s[1:len(s)]
		}
		if len(s) == 1 {
			fmt.Println()
			continue
		}
		t := []string{s[0]}
		for i := 1; i < len(s) && s[i] != s[0]; i++ {
			t = append(t, s[i])
		}
		fmt.Println(strings.Join(t, " "))
	}
}
