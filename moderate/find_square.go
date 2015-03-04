package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func dist2(p1, p2 []int) int {
	x, y := p2[0]-p1[0], p2[1]-p1[1]
	return x*x + y*y
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), ", ")
		u := make([][]int, 4)
		for ix, i := range t {
			var x, y int
			fmt.Sscanf(i, "(%d,%d)", &x, &y)
			u[ix] = []int{x, y}
		}
		d := []int{}
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				d = append(d, dist2(u[i], u[j]))
			}
		}
		sort.Ints(d)
		if d[0] == d[1] && d[0] == d[2] && d[0] == d[3] &&
			d[4] == d[5] && 2*d[0] == d[4] {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
