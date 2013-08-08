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
		t, p := strings.TrimSpace(s), 1
		for ix := 1; ix < len(t); ix++ {
			if t[ix] != t[ix-p] {
				p = ix + 1
			}
		}
		fmt.Println(p)
	}
}
