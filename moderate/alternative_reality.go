package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var m map[int]int

func init() {
	m = map[int]int{50: 25, 25: 10, 10: 5, 5: 1}
}

func alter(n, c int) int {
	if n == 0 || c == 1 {
		return 1
	}
	if c > n {
		return alter(n, m[c])
	}
	return alter(n-c, c) + alter(n, m[c])
}

func main() {
	var n int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &n)
		fmt.Println(alter(n, 50))
	}
}
