package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ") (")
		p, q := strings.Split(s[0][1:len(s[0])], ","), strings.Split(s[1][0:len(s[1])-1], ",")
		if len(p) < 2 || len(q) < 2 {
			fmt.Println(0)
			continue
		}
		var (
			st, av   []float32
			intersec int
		)
		for _, i := range p {
			var v float32
			fmt.Sscanf(i, "%f", &v)
			st = append(st, v)
		}
		for _, i := range q {
			var v float32
			fmt.Sscanf(i, "%f", &v)
			av = append(av, v)
		}
		if st[0] != 0 {
			st0 := st[0]
			for ix, i := range st {
				st[ix] = i - st0
			}
		}
		if av[0] != 0 {
			av0 := av[0]
			for ix, i := range av {
				av[ix] = i - av0
			}
		}
		stl, avl := st[len(st)-1], av[len(av)-1]
		for ix, i := range av {
			av[ix] = i * stl / avl
		}
		for i, j := 0, 0; i < len(st) && j < len(av); {
			if st[i] == av[j] {
				intersec++
				i++
				j++
			} else if st[i] < av[j] {
				i++
			} else if st[i] > av[j] {
				j++
			}
		}
		fmt.Println(len(st) + len(av) - intersec - 1)
	}
}
