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
		var m int
		fmt.Sscan(t[len(t)-1], &m)
		if m < len(t) {
			fmt.Println(t[len(t)-1-m])
		}
	}
}
