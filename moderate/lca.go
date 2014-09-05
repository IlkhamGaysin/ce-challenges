package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	parent := map[int]int{3: 8, 8: 30, 10: 20, 20: 8, 29: 20, 30: 0, 52: 30}
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var n1, n2 int
		fmt.Sscanf(scanner.Text(), "%d %d", &n1, &n2)
		for found := false; n1 != 0 && n1 != n2; n1 = parent[n1] {
			for i := n2; i != 0; i = parent[i] {
				if n1 == i {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		fmt.Println(n1)
	}
}
