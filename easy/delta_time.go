package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var h1, m1, s1, h2, m2, s2 int
		fmt.Sscanf(scanner.Text(), "%d:%d:%d %d:%d:%d", &h1, &m1, &s1, &h2, &m2, &s2)
		t := abs((h1-h2)*3600 + (m1-m2)*60 + s1 - s2)
		fmt.Printf("%02d:%02d:%02d\n", t/3600, (t/60)%60, t%60)
	}
}
