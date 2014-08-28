package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxis(a []int) (ret int) {
	ret = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > ret {
			ret = a[i]
		}
	}
	return ret
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	r := []int{}
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		t := make([]int, len(s))
		for ix, i := range s {
			fmt.Sscanf(i, "%d", &t[ix])
		}
		if len(r) > 0 {
			t[0] += r[0]
			if len(r) > 1 {
				t[len(r)] += r[len(r)-1]
			}
		}
		for i := 1; i < len(t)-1; i++ {
			t[i] += maxi(r[i-1], r[i])
		}
		r = make([]int, len(t))
		copy(r, t)
	}
	fmt.Println(maxis(r))
}
