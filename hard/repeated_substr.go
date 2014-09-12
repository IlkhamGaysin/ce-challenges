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
		l, m := len(scanner.Text()), ""
		if l > 0 {
			for n, c := 1, 0; n <= (l-c)/2; n++ {
				var f bool
				for i := c; i < l-n; i++ {
					subs := scanner.Text()[i : i+n]
					if len(strings.TrimSpace(subs)) == 0 {
						continue
					}
					if strings.Contains(scanner.Text()[i+n:l], subs) {
						m, c, f = subs, i, true
						break
					}
				}
				if !f {
					break
				}
			}
		}
		if m == "" {
			fmt.Println("NONE")
		} else {
			fmt.Println(m)
		}
	}
}
