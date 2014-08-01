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
	reader := bufio.NewReader(data)
	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		var n, m int
		fmt.Sscanf(string(s), "%d %d", &n, &m)
		if m == 0 {
			fmt.Println(n)
			continue
		} else if m == 1 {
			fmt.Println(n - 1)
			continue
		}
		l := make([]bool, n)
		for i := 0; i <= m%2; i++ {
			for j := 1; j < n; j += 2 {
				l[j] = true
			}
			for j := 2; j < n; j += 3 {
				l[j] = !l[j]
			}
		}
		l[n-1] = !l[n-1]
		var count int
		for i := 0; i < n; i++ {
			if !l[i] {
				count++
			}
		}
		fmt.Println(count)
	}
}
