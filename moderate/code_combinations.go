package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		var r int
		s := strings.Split(scanner.Text(), " | ")
		for i := 1; i < len(s); i++ {
			for j := 1; j < len(s[0]); j++ {
				code := []string{string(s[i][j]), string(s[i][j-1]), string(s[i-1][j]), string(s[i-1][j-1])}
				sort.Strings(code)
				if strings.Join(code, "") == "cdeo" {
					r += 1
				}
			}
		}
		fmt.Println(r)
	}
}
