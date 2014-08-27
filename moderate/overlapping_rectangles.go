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

func within(a []int, b point) bool {
	return b.x >= a[0] && b.x <= a[2] && b.y >= a[3] && b.y <= a[1]
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		t := make([]int, 8)
		for ix, i := range s {
			fmt.Sscanf(i, "%d", &t[ix])
		}
		p1 := []point{point{t[0], t[1]}, point{t[0], t[3]},
			point{t[2], t[1]}, point{t[2], t[3]}}
		p2 := []point{point{t[4], t[5]}, point{t[4], t[7]},
			point{t[6], t[5]}, point{t[6], t[7]}}
		var f bool
		for _, i := range p1 {
			if within(t[4:8], i) {
				f = true
				break
			}
		}
		if !f {
			for _, i := range p2 {
				if within(t[0:4], i) {
					f = true
					break
				}
			}
		}
		if f {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
