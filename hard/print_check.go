package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var s0, s1, s2, s3 []string

func wrd(a1, a2, a3 byte) bool {
	if a1+a2+a3 == 0 {
		return false
	}
	if a1 > 0 {
		fmt.Print(s0[a1] + s3[0])
	}
	if a2 > 1 {
		fmt.Print(s2[a2-2] + s0[a3])
	} else if a2 > 0 {
		fmt.Print(s1[a3])
	} else {
		fmt.Print(s0[a3])
	}
	return true
}

func main() {
	s0 = []string{"", "One", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine"}
	s1 = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
		"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	s2 = []string{"Twenty", "Thirty", "Forty", "Fifty",
		"Sixty", "Seventy", "Eighty", "Ninety"}
	s3 = []string{"Hundred", "Thousand", "Million"}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var b int
		fmt.Sscanf(scanner.Text(), "%d", &b)
		c := make([]byte, 9)
		if b == 0 {
			fmt.Print("Zero")
		} else {
			for i := 8; b > 0 && i >= 0; i-- {
				c[i], b = byte(b%10), b/10
			}
			if wrd(c[0], c[1], c[2]) {
				fmt.Print(s3[2])
			}
			if wrd(c[3], c[4], c[5]) {
				fmt.Print(s3[1])
			}
			_ = wrd(c[6], c[7], c[8])
		}
		fmt.Println("Dollars")
	}
}
