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

	digits := map[string]int{"zero": 0, "one": 1, "two": 2,
		"three": 3, "four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9}

	reader := bufio.NewReader(data)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Split(strings.TrimSpace(s), ";")
		for _, i := range t {
			fmt.Print(digits[i])
		}
		fmt.Println()
	}
}
