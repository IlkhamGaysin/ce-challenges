package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const morse = "ETIANMSURWDKGOHVF L PJBXCYZQ  54 3   2       16       7   8 90"

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	m := 1
	for scanner.Scan() {
		c := rune(scanner.Text()[0])
		switch {
		case c == '.':
			m <<= 1
		case c == '-':
			m = (m << 1) + 1
		case m == 1:
			fmt.Print(" ")
		default:
			if m < 64 {
				fmt.Print(string(morse[m-2]))
			}
			if c == '\n' {
				fmt.Println()
			}
			m = 1
		}
	}
	if m > 1 {
		if m < 64 {
			fmt.Print(string(morse[m-2]))
		}
		fmt.Println()
	}
}
