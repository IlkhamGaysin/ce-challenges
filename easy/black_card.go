package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var m int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " | ")
		fmt.Sscanf(t[1], "%d", &m)
		u := strings.Fields(t[0])
		for len(u) > 1 {
			n := m%len(u) - 1
			if n == -1 {
				u = u[:len(u)-1]
			} else {
				u = append(u[:n], u[n+1:]...)
			}
		}
		fmt.Println(u[0])
	}
}
