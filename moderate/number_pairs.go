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
	reader := bufio.NewReader(data)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		t := strings.Split(strings.TrimSpace(s), ";")
		u := strings.Split(t[0], ",")
		var n int
		fmt.Sscan(t[1], &n)
		v, vk := make(map[int]bool, len(u)), []int{}
		for i := range u {
			var w int
			fmt.Sscan(u[i], &w)
			v[w], vk = true, append(vk, w)
		}
		f := false
		for _, i := range vk {
			if 2*i < n && v[n-i] {
				if f {
					fmt.Print(";")
				}
				fmt.Print(i, ",", n-i)
				f = true
			}
		}
		if !f {
			fmt.Print("NULL")
		}
		fmt.Println()
	}
}
