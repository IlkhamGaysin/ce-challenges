package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		m, mink, minv := make(map[int]int), math.MaxInt32, 0
		for ix, i := range strings.Fields(strings.TrimSpace(s)) {
			k, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
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
