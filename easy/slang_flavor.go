package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	phrases := []string{", yeah!", ", this is crazy, I tell ya.",
		", can U believe this?", ", eh?", ", aw yea.",
		", yo.", "? No way!", ". Awesome!"}
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	var c int
	var l bool
	for scanner.Scan() {
		t := scanner.Text()
		if t == "." || t == "!" || t == "?" {
			if l {
				fmt.Print(phrases[c])
				l, c = false, (c+1)%len(phrases)
			} else {
				fmt.Print(t)
				l = true
			}
		} else {
			fmt.Print(t)
		}
	}
}
