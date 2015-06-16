package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var m int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var ret []string
		t := strings.Split(scanner.Text(), ";")
		fmt.Sscan(t[1], &m)
		u := strings.Split(t[0], ",")
		mi := m
		for ix := range u {
			var o int
			if mi <= len(u)-ix {
				mi--
				o = mi - ix%m
				if mi == 0 {
					mi = m
				}
			}
			ret = append(ret, u[ix+o])
		}
		fmt.Println(strings.Join(ret, ","))
	}
}
