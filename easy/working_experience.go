package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var moy map[string]int

func month(s string) int {
	k, err := strconv.Atoi(s[4:8])
	if err != nil {
		log.Fatal(err)
	}
	return 12*(k-1990) + moy[s[0:3]]
}

func main() {
	moy = map[string]int{"Jan": 0, "Feb": 1, "Mar": 2, "Apr": 3,
		"May": 4, "Jun": 5, "Jul": 6, "Aug": 7,
		"Sep": 8, "Oct": 9, "Nov": 10, "Dec": 11}

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
		work := make(map[int]bool)
		t := strings.Split(s, "; ")
		for _, i := range t {
			t0, t1 := month(i[0:8]), month(i[9:17])
			for j := t0; j <= t1; j++ {
				work[j] = true
			}
		}
		fmt.Println(len(work) / 12)
	}
}
