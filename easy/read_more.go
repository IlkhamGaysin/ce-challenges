package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 55 {
			for i := 40; i > 0; i-- {
				if line[i] == ' ' {
					line = line[:i]
					break
				}
			}
			if len(line) > 55 {
				line = line[:40]
			}
			line += "... <Read More>"
		}
		fmt.Println(line)
	}
}
