package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func recovery(s, t string) string {
	var k int
	words, sequ := strings.Fields(s), strings.Fields(t)
	r := make([]string, len(words))
	for ix, i := range sequ {
		fmt.Sscan(i, &k)
		r[k-1] = words[ix]
	}
	for ix := 0; ix < len(r); ix++ {
		if r[ix] == "" {
			r[ix] = words[len(words)-1]
			break
		}
	}
	return strings.Join(r, " ")
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), ";")
		if len(t) == 2 {
			fmt.Println(recovery(t[0], t[1]))
		}
	}
}
