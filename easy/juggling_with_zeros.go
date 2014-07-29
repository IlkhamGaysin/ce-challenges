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
		for i := 0; i < len(t)/2; i++ {
			num <<= uint(len(t[2*i+1]))
			if t[2*i] == "00" {
				for j := 0; j < len(t[2*i+1]); j++ {
					num += 1 << uint(j)
				}
			}
		}
		fmt.Println(num)
	}
}
