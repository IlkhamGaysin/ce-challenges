package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var sx, tx int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		s2, t2 := strings.Fields(strings.TrimSpace(s[0])), strings.Fields(strings.TrimSpace(s[1]))
		var r []string
		for i := range t2 {
			fmt.Sscan(s2[i], &sx)
			fmt.Sscan(t2[i], &tx)
			r = append(r, fmt.Sprint(sx*tx))
		}
		fmt.Println(strings.Join(r, " "))
	}
}
