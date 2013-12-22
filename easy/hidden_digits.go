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
	reader := bufio.NewReader(data)
	var s string
	for {
		s, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		maph := func(r rune) rune {
			if strings.Contains("0123456789", string(r)) {
				return r
			} else if strings.Contains("abcdefghij", string(r)) {
				return r - 'a' + '0'
			}
			return -1
		}
		t := strings.Map(maph, s)

		if len(t) == 0 {
			fmt.Println("NONE")
		} else {
			fmt.Println(t)
		}
	}
}
