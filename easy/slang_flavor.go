package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		phrases = []string{", yeah!", ", this is crazy, I tell ya.",
			", can U believe this?", ", eh?", ", aw yea.",
			", yo.", "? No way!", ". Awesome!"}
		c int
		l bool
		t byte = '\n'
	)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		t = scanner.Text()[0]
		if t == '.' || t == '!' || t == '?' {
			l = !l
			if !l {
				fmt.Print(phrases[c])
				c = (c + 1) % len(phrases)
				continue
			}
		}
		fmt.Printf("%c", t)
	}
	if t != '\n' {
		fmt.Println()
	}
}
