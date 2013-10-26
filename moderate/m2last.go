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
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Fields(strings.TrimSpace(s))
		var m int
		fmt.Sscan(t[len(t)-1], &m)
		if m < len(t) {
			fmt.Println(t[len(t)-1-m])
		}
	}
}
