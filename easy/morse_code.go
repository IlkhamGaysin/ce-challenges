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
	reader := bufio.NewReader(data)
	m := 1
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}
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
}
