package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func numzero(t []int, c, z int) int {
	switch {
	case c == 0:
		if z == 0 {
			return 1
		} else {
			return 0
		}
	case len(t) < c || z+c*t[0] > 0 || z+c*t[len(t)-1] < 0:
		return 0
	default:
		return numzero(t[1:len(t)], c-1, z+t[0]) + numzero(t[1:len(t)], c, z)
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		t := make([]int, len(s))
		for ix, i := range s {
			fmt.Sscanf(i, "%d", &t[ix])
		}
		sort.Ints(t)
		fmt.Println(numzero(t, 4, 0))
	}
}
