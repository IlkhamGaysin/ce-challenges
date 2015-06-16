package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var n, k int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var v []int
		t := strings.Split(scanner.Text(), " | ")
		fmt.Sscanf(t[1], "%d", &n)
		u := strings.Fields(t[0])
		for _, i := range u {
			fmt.Sscanf(i, "%d", &k)
			v = append(v, k)
		}
		for i := 0; i < len(v)-n; i++ {
			sort.Ints(v[i : i+n+1])
		}
		for ix, i := range v {
			u[ix] = fmt.Sprint(i)
		}
		fmt.Println(strings.Join(u, " "))
	}
}
