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
		t := strings.Fields(scanner.Text())
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
