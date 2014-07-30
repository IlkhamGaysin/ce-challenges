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
		var num int
		t := strings.Fields(strings.TrimSpace(s))
		for i := 0; i < len(t); i += 2 {
			num <<= uint(len(t[i+1]))
			if t[i] == "00" {
				num += (1 << uint(len(t[i+1]))) - 1
			}
		}
		fmt.Println(num)
	}
}
