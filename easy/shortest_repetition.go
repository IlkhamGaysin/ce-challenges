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
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t, p := scanner.Text(), 1
		for ix := 1; ix < len(t); ix++ {
			if t[ix] != t[ix-p] {
				p = ix + 1
			}
		}
		fmt.Println(p)
	}
}
