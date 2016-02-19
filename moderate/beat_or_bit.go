package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func g2d(a uint) uint {
	a ^= a >> 4
	a ^= a >> 2
	a ^= a >> 1
	return a
}

func main() {
	var v uint
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")
		for i := 0; i < len(s); i++ {
			fmt.Sscanf(s[i], "%b", &v)
			s[i] = fmt.Sprint(g2d(v))
		}
		fmt.Println(strings.Join(s, " | "))
	}
}
