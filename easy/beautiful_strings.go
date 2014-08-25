package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func clean(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char
	}
	return -1
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		c := make([]int, 26)
		t := strings.Map(clean, strings.ToLower(string(s)))
		for _, i := range t {
			c[i-'a']++
		}
		sort.Ints(c)
		r := 0
		for i := 26; i > 0 && c[i-1] > 0; i-- {
			r += i * c[i-1]
		}
		fmt.Println(r)
	}
}
