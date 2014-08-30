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
		var f bool
		var next int
		for i := strings.Index(s[1], string(s[0][0])); i >= 0; i += next {
			if strings.HasSuffix(s[1], string(s[0][0:len(s[0])-i])) && (i == 0 || strings.HasPrefix(s[1], string(s[0][len(s[0])-i:len(s[0])]))) {
				f = true
				break
			}
			next = 1 + strings.Index(s[1][i+1:len(s[1])], string(s[0][0]))
			if next == 0 {
				break
			}
		}
		if f {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
