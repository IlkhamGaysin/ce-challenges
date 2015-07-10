package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func inLine(a, b, c point) bool {
	dx1, dy1, dx2, dy2 := a.x-b.x, a.y-b.y, a.x-c.x, a.y-c.y
	return dx1*dy2 == dx2*dy1
}

func line(a, b int, t []point) (l []int) {
	for i := 0; i < len(t); i++ {
		if i == a || i == b {
			l = append(l, i)
			continue
		}
		if inLine(t[i], t[a], t[b]) {
			if i < b {
				return l
			}
			l = append(l, i)
		}
	}
	return l
}

func main() {
	var x, y int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")
		var t []point
		for _, i := range s {
			fmt.Sscanf(i, "%d %d", &x, &y)
			t = append(t, point{x, y})
		}
		var r [][]int
		for i := 0; i < len(t)-2; i++ {
			for j := i + 1; j < len(t)-1; j++ {
				if l := line(i, j, t); len(l) > 2 {
					r = append(r, l)
				}
			}
		}
		fmt.Println(len(r))
	}
}
