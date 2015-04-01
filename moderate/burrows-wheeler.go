package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		s := strings.TrimRight(scanner.Text(), "|")
		e := make([]int, len(s))
		var k int
		for i := 0; i < len(s); i++ {
			e[i] = int(s[i])<<16 + i
			if s[i] == '$' {
				k = i
			}
		}
		sort.Ints(e)
		for i := 0; i < len(s); i, k = i+1, e[k]%(1<<16) {
			fmt.Printf("%c", e[k]>>16)
		}
		fmt.Println()
	}
}
