package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var t string
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var r int
		s := strings.Split(scanner.Text(), " | ")
		for i := 0; i < len(s[0]); i++ {
			if s[0][i] != s[1][i] {
				r++
				if r > 6 {
					break
				}
			}
		}
		switch {
		case r == 0:
			t = "Done"
		case r <= 2:
			t = "Low"
		case r <= 4:
			t = "Medium"
		case r <= 6:
			t = "High"
		default:
			t = "Critical"
		}
		fmt.Println(t)
	}
}
