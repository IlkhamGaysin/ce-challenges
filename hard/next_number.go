package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var r []int
func dsig(a int) (ret int) {
	r = r[:0]
	for a > 0 {
		if a%10 > 0 {
			r = append(r, a%10)
		}
		a /= 10
	}
	sort.Ints(r)
	for _, i := range r {
		ret = 10*ret + i
	}
	return ret
}

func main() {
	r = []int{}
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var d int
		fmt.Sscanf(scanner.Text(), "%d", &d)
		m9, ds, e := d%9, dsig(d), d+1
		for e%9 != m9 || dsig(e) != ds {
			e++
		}
		fmt.Println(e)
	}
}
