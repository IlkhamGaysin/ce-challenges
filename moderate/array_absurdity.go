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
		var n int
		fmt.Sscanf(s[0], "%d", &n)
		t := strings.Split(s[1], ",")
		m := make(map[int]bool, n-1)
		for _, i := range t {
			var v int
			fmt.Sscanf(i, "%d", &v)
			if m[v] {
				fmt.Println(i)
				break
			}
			m[v] = true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
