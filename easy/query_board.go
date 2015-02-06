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
	m := make([][]int, 256)
	for i := 0; i < 256; i++ {
		m[i] = make([]int, 256)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Fields(scanner.Text())
		if strings.HasPrefix(t[0], "Query") {
			var p, sum int
			fmt.Sscanf(t[1], "%d", &p)
			if strings.HasSuffix(t[0], "Row") {
				for i := 0; i < 256; i++ {
					sum += m[p][i]
				}
			} else {
				for i := 0; i < 256; i++ {
					sum += m[i][p]
				}
			}
			fmt.Println(sum)
		} else {
			var p, v int
			fmt.Sscanf(t[1], "%d", &p)
			fmt.Sscanf(t[2], "%d", &v)
			if strings.HasSuffix(t[0], "Row") {
				for i := 0; i < 256; i++ {
					m[p][i] = v
				}
			} else {
				for i := 0; i < 256; i++ {
					m[i][p] = v
				}
			}
		}
	}
}
