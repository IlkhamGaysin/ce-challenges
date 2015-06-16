package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var v int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var (
			c int
			l []int
			m []string
		)
		t := strings.Split(scanner.Text(), ";")
		for _, i := range t {
			if u := strings.Split(i, ","); len(u) > 1 {
				_, err := fmt.Sscan(u[1], &v)
				if err != nil {
					log.Fatal(err)
				}
				l = append(l, v)
			}
		}
		sort.Ints(l)
		for _, i := range l {
			m = append(m, fmt.Sprint(i-c))
			c = i
		}
		fmt.Println(strings.Join(m, ","))
	}
}
