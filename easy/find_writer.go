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
		} else if len(s) < 3 {
			continue
		}
		t := strings.Split(s, "|")
		u := strings.Fields(strings.TrimSpace(t[1]))
		for _, i := range u {
			var f int
			fmt.Sscan(i, &f)
			fmt.Print(string(t[0][f-1]))
		}
		fmt.Println()
	}
}
