package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func diagonal(c, h []int) bool {
	a, b, p, q := float64(h[0]), float64(h[1]), float64(c[0]), float64(c[1])
	return (p > a) && (b >= (2*p*q*a+(p*p-q*q)*math.Sqrt(p*p+q*q-a*a))/(p*p+q*q))
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	h, b := make([]int, 2), make([]int, 3)
	var s1, s2, s3, s4, s5, s6 int
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "|")
		u := strings.Split(t[1], ";")
		fmt.Sscanf(t[0], "[%d,%d] [%d,%d]", &s1, &s2, &s3, &s4)
		h[0], h[1] = abs(s1-s3), abs(s2-s4)
		sort.Ints(h)
		var r []int
		for _, i := range u {
			var j int
			fmt.Sscanf(i, "(%d [%d,%d,%d] [%d,%d,%d])", &j, &s1, &s2, &s3, &s4, &s5, &s6)
			b[0], b[1], b[2] = abs(s1-s4), abs(s2-s5), abs(s3-s6)
			sort.Ints(b)
			if b[0] <= h[0] && b[1] <= h[1] { // || diagonal(b, h)
				r = append(r, j)
			}
		}
		if len(r) == 0 {
			fmt.Println("-")
		} else {
			sort.Ints(r)
			var r2 []string
			for _, i := range r {
				r2 = append(r2, fmt.Sprint(i))
			}
			fmt.Println(strings.Join(r2, ","))
		}
	}
}
