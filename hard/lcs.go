package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func bt(c [][]int, x, y string, i, j int) string {
	if i < 1 || j < 1 {
		return ""
	}
	if x[i-1] == y[j-1] {
		return bt(c, x, y, i-1, j-1) + string(x[i-1])
	}
	if c[i][j-1] > c[i-1][j] {
		return bt(c, x, y, i, j-1)
	}
	return bt(c, x, y, i-1, j)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		} else if len(s) < 3 {
			continue
		}
		t := strings.Split(strings.TrimSpace(s), ";")
		c := make([][]int, len(t[0])+1)
		c[0] = make([]int, len(t[1])+1)
		for ix, i := range t[0] {
			c[ix+1] = make([]int, len(t[1])+1)
			for jx, j := range t[1] {
				if i == j {
					c[ix+1][jx+1] = c[ix][jx] + 1
				} else {
					c[ix+1][jx+1] = max(c[ix+1][jx], c[ix][jx+1])
				}
			}
		}
		fmt.Println(bt(c, t[0], t[1], len(t[0]), len(t[1])))
	}
}
