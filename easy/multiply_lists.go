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
		s, err = reader.ReadString('|')
		if err != nil {
			break
		}
		t, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s2, t2 := strings.Fields(strings.TrimSpace(s)), strings.Fields(strings.TrimSpace(t))
		var r []string
		for i := 0; i < len(t2); i++ {
			sx, err := strconv.Atoi(s2[i])
			if err != nil {
				log.Fatal(err)
			}
			tx, err := strconv.Atoi(t2[i])
			if err != nil {
				log.Fatal(err)
			}
			r = append(r, fmt.Sprint(sx*tx))
		}
		fmt.Println(strings.Join(r, " "))
	}
}
