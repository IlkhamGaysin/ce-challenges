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
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var n int
	fmt.Sscan(strings.TrimSpace(s), &n)
	lol := make([]string, n)
	for {
		s, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		for ix, i := range lol {
			if len(s) > len(i) {
				lol = append(lol[:ix], append([]string{s}, lol[ix:n-1]...)...)
				break
			}
		}
	}
	for _, i := range lol {
		fmt.Print(i)
	}
}
