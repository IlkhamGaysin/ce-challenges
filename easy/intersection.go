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
		t := strings.Split(scanner.Text(), ";")
		u, v := strings.Split(t[0], ","), strings.Split(t[1], ",")
		x, y := []int{}, []int{}
		for _, i := range u {
			var a int
			fmt.Sscanf(i, "%d", &a)
			x = append(x, a)
		}
		for _, i := range v {
			var a int
			fmt.Sscanf(i, "%d", &a)
			y = append(y, a)
		}
		r := []string{}
		for i, j := 0, 0; i < len(u) && j < len(v); {
			if x[i] == y[j] {
				r = append(r, fmt.Sprint(x[i]))
				i++
				j++
			} else if x[i] > y[j] {
				j++
			} else {
				i++
			}
		}
		fmt.Println(strings.Join(r, ","))
	}
}
