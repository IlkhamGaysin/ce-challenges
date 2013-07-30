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
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var n int
		fmt.Sscan(s, &n)
		l, st := []int{1}, []int{1}
		for i := 0; i < n-1; i++ {
			m := append([]int{0}, l...)
			for jx, j := range l {
				m[jx] += j
			}
			st = append(st, m...)
			l = m
		}
		t := fmt.Sprint(st)
		fmt.Println(t[1 : len(t)-1])
	}
}
