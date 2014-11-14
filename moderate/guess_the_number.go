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
	scanner.Split(bufio.ScanWords)
	lo, hi := 0, -1
	for scanner.Scan() {
		switch s, c := scanner.Text(), (lo+hi)%2; {
		case hi < 0:
			fmt.Sscan(s, &hi)
		case s == "Lower":
			hi = (lo+hi)/2 + c - 1
		case s == "Higher":
			lo = (lo+hi)/2 + c + 1
		case s == "Yay!":
			fmt.Println((lo+hi)/2 + c)
			lo, hi = 0, -1
		}
	}
}
