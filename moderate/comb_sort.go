package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const shrink = 1.25

func main() {
	var (
		n, k, g, l int
		c          bool
	)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		u := strings.Fields(scanner.Text())
		n, g, l, c = 0, len(u), 0, false
		v := make([]int, len(u))
		for ix, i := range u {
			fmt.Sscanf(i, "%d", &k)
			v[ix] = k
		}
		for g > 1 || c {
			n, c = n+1, false
			if g > 1 {
				g = int(float32(g) / shrink)
			}
			for j := 0; j < len(u)-g; j++ {
				if u[j] > u[j+g] {
					u[j], u[j+g], c, l = u[j+g], u[j], true, n
				}
			}
		}
		fmt.Println(l)
	}
}
