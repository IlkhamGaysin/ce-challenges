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
		s := strings.Split(scanner.Text(), ",")
		var l, m int
		fmt.Sscanf(strings.TrimSpace(s[0]), "%d", &m)
		for _, i := range s {
			var v int
			fmt.Sscanf(strings.TrimSpace(i), "%d", &v)
			if v+l > m {
				m = v + l
			}
			if v+l > 0 {
				l += v
			} else {
				l = 0
			}
		}
		fmt.Println(m)
	}
}
