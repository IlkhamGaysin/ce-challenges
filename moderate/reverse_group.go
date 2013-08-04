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
		var m int
		fmt.Sscan(t[1], &m)
		u := strings.Split(t[0], ",")
		ret := []string{}
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
