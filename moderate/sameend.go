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
		switch t := strings.Split(scanner.Text(), ","); {
		case len(t) < 2:
		case strings.HasSuffix(t[0], t[1]):
			fmt.Println(1)
		default:
			fmt.Println(0)
		}
	}
}
