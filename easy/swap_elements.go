package main

import (
	"bufio"
	"fmt"
	"log"
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
	var s, t string
	for {
		s, err = reader.ReadString(':')
		if err != nil {
			break
		}
		t, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s2, t2 := strings.Fields(strings.TrimSpace(s)), strings.Fields(strings.TrimSpace(t))
		r := make([]string, len(s2)-1)
		copy(r, s2)
		for _, i := range t2 {
			t3 := strings.Split(i, "-")
			sw1, err := strconv.Atoi(t3[0])
			if err != nil {
				log.Fatal(err)
			}
			sw2, err := strconv.Atoi(strings.TrimRight(t3[1], ","))
			if err != nil {
				log.Fatal(err)
			}
			r[sw1], r[sw2] = r[sw2], r[sw1]
		}
		fmt.Println(strings.Join(r, " "))
	}
}
