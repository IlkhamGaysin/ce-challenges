package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var n int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), ";")
		u := strings.Split(t[0], ",")
		fmt.Sscan(t[1], &n)
		v, vk := make(map[int]bool, len(u)), make([]int, len(u))
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
