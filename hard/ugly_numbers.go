package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func pow3(a int) int {
	ret := 1
	for i := 0; i < a; i++ {
		ret *= 3
	}
	return ret
}

func ugly(j int, n []int) bool {
	var k int
	var s int64
	c, cj, p := int64(n[0]), j, true
	for k < len(n)-1 {
		ops := cj % 3
		cj /= 3
		if ops == 0 {
			c *= 10
		} else {
			if p {
				s += c
				if ops == 1 {
					p = false
				}
			} else {
				s -= c
				if ops == 2 {
					p = true
				}
			}
			c = 0
		}
		k++
		c += int64(n[k])
	}
	if p {
		s += c
	} else {
		s -= c
	}
	if s%2 == 0 || s%3 == 0 || s%5 == 0 || s%7 == 0 {
		return true
	}
	return false
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var u int
		t := []int{}
		for _, i := range scanner.Text() {
			t = append(t, int(i-'0'))
		}
		p := pow3(len(t) - 1)
		for j := 0; j < p; j++ {
			if ugly(j, t) {
				u++
			}
		}
		fmt.Println(u)
	}
}
