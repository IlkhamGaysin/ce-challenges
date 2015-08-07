package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	var a int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var (
			r []int
			u []string
		)
		s := strings.Split(scanner.Text(), "|")
		for ix, i := range s {
			t := strings.Fields(strings.TrimSpace(i))
			for jx, j := range t {
				fmt.Sscanf(j, "%d", &a)
				if ix == 0 {
					r = append(r, a)
				} else {
					r[jx] = max(r[jx], a)
				}
			}
		}
		for _, i := range r {
			u = append(u, fmt.Sprint(i))
		}
		fmt.Println(strings.Join(u, " "))
	}
}
