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
		t = '\n'
	)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		t = rune(scanner.Text()[0])
		if t == '.' || t == '!' || t == '?' {
			if l {
				fmt.Print(phrases[c])
				l, c = false, (c+1)%len(phrases)
			} else {
				fmt.Printf("%c", t)
				l = true
			}
		} else {
			fmt.Printf("%c", t)
		}
	}
	if t != '\n' {
		fmt.Println()
	}
}
