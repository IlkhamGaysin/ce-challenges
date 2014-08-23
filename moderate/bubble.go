package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		t := strings.Split(strings.TrimSpace(string(s)), " | ")
		n, err := strconv.Atoi(t[1])
		if err != nil {
			log.Fatal(err)
		}
		u := strings.Fields(t[0])
		v := []int{}
		for _, i := range u {
			k, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
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
