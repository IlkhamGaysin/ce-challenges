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
		m := len(s[0])
		for _, i := range s {
			if strings.Contains(i, "XY") {
				m = 0
				continue
			}
			n := strings.Count(i, ".")
			if n < m {
				m = n
			}
		}
		fmt.Println(m)
	}
}
