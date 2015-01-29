package main

import (
	"bufio"
	"fmt"
	"log"
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
		s := strings.Split(scanner.Text(), "|")
		s2, t2 := strings.Fields(strings.TrimSpace(s[0])), strings.Fields(strings.TrimSpace(s[1]))
		var r []string
		for i := 0; i < len(t2); i++ {
			var sx, tx int
			fmt.Sscanf(s2[i], "%d", &sx)
			fmt.Sscanf(t2[i], "%d", &tx)
			r = append(r, fmt.Sprint(sx*tx))
		}
		fmt.Println(strings.Join(r, " "))
	}
}
