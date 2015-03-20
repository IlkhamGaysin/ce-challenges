package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func contains(a []int, b int) bool {
	ix := sort.SearchInts(a, b)
	return ix < len(a) && a[ix] == b
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		var n, e, r int
		fmt.Sscanf(s[0], "%d %d", &n, &e)
		g := make(map[int][]int)
		t := strings.Split(s[1], ",")
		p := make([]int, 2)
		for _, i := range t {
			fmt.Sscanf(i, "%d %d", &p[0], &p[1])
			sort.Ints(p)
			g[p[0]] = append(g[p[0]], p[1])
		}
		for k := range g {
			sort.Ints(g[k])
		}
		for _, v := range g {
			for i := 0; i < len(v)-1; i++ {
				for j := i + 1; j < len(v); j++ {
					if contains(g[v[i]], v[j]) {
						r++
					}
				}
			}
		}
		fmt.Println(r)
	}
}
