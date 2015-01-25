package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		m, mink, minv := make(map[int]int), math.MaxInt32, 0
		for ix, i := range strings.Fields(strings.TrimSpace(scanner.Text())) {
			var k int
			fmt.Sscanf(i, "%d", &k)
			if m[k] == 0 {
				m[k] = ix + 1
			} else {
				m[k] = math.MaxInt32
			}
		}
		for k, v := range m {
			if v < math.MaxInt32 && k < mink {
				mink, minv = k, v
			}
		}
		if mink == math.MaxInt32 {
			fmt.Println(0)
		} else {
			fmt.Println(minv)
		}
	}
}
