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
		s := strings.Split(scanner.Text(), ":")
		r, t := strings.Fields(strings.TrimSpace(s[0])), strings.Fields(strings.TrimSpace(s[1]))
		for _, i := range t {
			var sw1, sw2 int
			fmt.Sscanf(i, "%d-%d,", &sw1, &sw2)
			r[sw1], r[sw2] = r[sw2], r[sw1]
		}
		fmt.Println(strings.Join(r, " "))
	}
}
