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
	reader := bufio.NewReader(data)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		} else if len(s) < 3 {
			continue
		}
		t := strings.Split(strings.TrimRight(s, "\n"), ",")
		if strings.HasSuffix(t[0], t[1]) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
