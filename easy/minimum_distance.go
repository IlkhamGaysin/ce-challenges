package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var n, r int
		s := strings.Fields(scanner.Text())
		fmt.Sscanf(s[0], "%d", &n)
		t := make([]int, n)
		for i := 1; i <= n; i++ {
			fmt.Sscanf(s[i], "%d", &t[i-1])
		}
		sort.Ints(t)
		for i, u := 0, t[n/2]; i < n; i++ {
			r += abs(t[i] - u)
		}
		fmt.Println(r)
	}
}
