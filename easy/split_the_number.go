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
	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		t := strings.Fields(string(s))
		var n, m, p int
		var neg bool
		if strings.Contains(t[1], "+") {
			p = strings.Index(t[1], "+")
		} else {
			p = strings.Index(t[1], "-")
			neg = true
		}
		fmt.Sscanf(t[0][0:p], "%d", &n)
		fmt.Sscanf(t[0][p:len(t[0])], "%d", &m)
		if neg {
			m = -m
		}
		fmt.Println(n + m)
	}
}
