package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		r := 1
		s := scanner.Text()
		n := len(s) / 2
		for i := 0; i < n; i++ {
			if ((s[i] == 'A') && (s[i+n] == 'B')) ||
			   ((s[i] == 'B') && (s[i+n] == 'A')) {
				r = 0
				break
			} else if (s[i] == '*') && (s[i+n] == '*') {
				r *= 2
			}
		}
		fmt.Println(r)
	}
}
