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
	c := make([]int, 26)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Map(clean, strings.ToLower(scanner.Text()))
		for _, i := range t {
			c[i-'a']++
		}
		sort.Ints(c)
		r := 0
		for i := 26; i > 0 && c[i-1] > 0; i-- {
			r += i * c[i-1]
			c[i-1] = 0
		}
		fmt.Println(r)
	}
}
