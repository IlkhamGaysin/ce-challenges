package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, k int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " | ")
		fmt.Sscanf(t[1], "%d", &n)
		u := strings.Fields(t[0])
		n = min(n, len(u)/2)
		v := make([]int, len(u))
		for ix, i := range u {
			fmt.Sscanf(i, "%d", &k)
			v[ix] = k
		}
		for i := 0; i < n; i++ {
			for j := i + 1; j < len(v)-i; j++ {
				if v[j-1] > v[j] {
					v[j-1], v[j] = v[j], v[j-1]
				}
			}
			for j := len(v) - i - 2; j >= i+1; j-- {
				if v[j-1] > v[j] {
					v[j-1], v[j] = v[j], v[j-1]
				}
			}
		}
		for ix, i := range v {
			u[ix] = fmt.Sprint(i)
		}
		fmt.Println(strings.Join(u, " "))
	}
}
