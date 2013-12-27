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
		}
		t := strings.Split(s, ";")
		l := []int{}
		for _, i := range t {
			u := strings.Split(i, ",")
			if len(u) > 1 {
				var v int
				_, err := fmt.Sscan(u[1], &v)
				if err != nil {
					log.Fatal(err)
				}
				l = append(l, v)
			}
		}
		sort.Ints(l)
		m := []string{}
		var c int
		for _, i := range l {
			m = append(m, fmt.Sprint(i-c))
			c = i
		}
		fmt.Println(strings.Join(m, ","))
	}
}
