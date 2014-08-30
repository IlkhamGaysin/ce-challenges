package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	stats := [][]int{[]int{2, 4, 8, 6}, []int{3, 9, 7, 1},
		[]int{4, 6}, []int{5}, []int{6},
		[]int{7, 9, 3, 1}, []int{8, 4, 2, 6},
		[]int{9, 1}}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var a, n int
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &n)
		a -= 2
		m := n / len(stats[a])
		res := make([]int, 10)
		for _, i := range stats[a] {
			res[i] += m
		}
		for i := 0; i < n%len(stats[a]); i++ {
			res[stats[a][i]]++
		}
		resu := make([]string, 10)
		for i := 0; i < 10; i++ {
			resu[i] = fmt.Sprintf("%d: %d", i, res[i])
		}
		fmt.Println(strings.Join(resu, ", "))
	}
}
