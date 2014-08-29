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
		var l, d, n, count, t int
		fmt.Sscanf(scanner.Text(), "%d %d %d", &l, &d, &n)
		s := strings.Fields(scanner.Text())
		tx := 6 - d
		for i := 6; i <= l-6; i += d {
			if i > tx-d {
				i = tx
				if t == n {
					tx = l - 6 + d
				} else {
					fmt.Sscanf(s[t+3], "%d", &tx)
					t++
				}
			} else {
				count++
			}
		}
		fmt.Println(count)
	}
}
