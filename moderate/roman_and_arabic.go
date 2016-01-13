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
	scanner.Split(bufio.ScanBytes)
	var num, ar, rr int
	for scanner.Scan() {
		c := scanner.Text()[0]
		if c == '\n' {
			fmt.Printf("%d\n", num+ar*rr)
			num, ar, rr = 0, 0, 0
			continue
		}
		var (
			r int
			a = int(c - '0')
		)
		switch scanner.Scan(); scanner.Text()[0] {
		case 'M':
			r = 1000
		case 'D':
			r = 500
		case 'C':
			r = 100
		case 'L':
			r = 50
		case 'X':
			r = 10
		case 'V':
			r = 5
		case 'I':
			r = 1
		}
		if r > rr {
			num -= ar * rr
		} else {
			num += ar * rr
		}
		ar, rr = a, r
	}
	if rr > 0 {
		fmt.Printf("%d\n", num+ar*rr)
	}
}
