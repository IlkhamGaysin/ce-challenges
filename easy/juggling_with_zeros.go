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
		var num uint64
		t := strings.Fields(scanner.Text())
		for i := 0; i < len(t); i += 2 {
			num <<= uint(len(t[i+1]))
			if t[i] == "00" {
				num += (1 << uint(len(t[i+1]))) - 1
			}
		}
		fmt.Println(num)
	}
}
