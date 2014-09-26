package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	digits = `
-**----*--***--***---*---****--**--****--**---**--
*--*--**-----*----*-*--*-*----*-------*-*--*-*--*-
*--*---*---**---**--****-***--***----*---**---***-
*--*---*--*-------*----*----*-*--*--*---*--*----*-
-**---***-****-***-----*-***---**---*----**---**--
--------------------------------------------------
`
	w, h = 5, 6
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		r := make([]string, h)
		for _, i := range scanner.Text() {
			if i >= '0' && i <= '9' {
				for j := 0; j < h; j++ {
					pos := 1 + j*w*10 + j + int(i-'0')*w
					r[j] += digits[pos : pos+w]
				}
			}
		}
		for i := 0; i < h; i++ {
			fmt.Println(r[i])
		}
	}
}
