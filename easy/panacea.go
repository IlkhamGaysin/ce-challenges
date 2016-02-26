package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var v int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var m int
		s := strings.Split(scanner.Text(), " | ")
		l := strings.Fields(s[0])
		r := strings.Fields(s[1])
		for _, i := range l {
			fmt.Sscanf(i, "%x", &v)
			m += v
		}
		for _, i := range r {
			fmt.Sscanf(i, "%b", &v)
			m -= v
		}
		if m > 0 {
			fmt.Println("False")
		} else {
			fmt.Println("True")
		}
	}
}
