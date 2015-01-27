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
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var t, r1, r2 []string
		t = strings.Split(scanner.Text(), ",")
		for _, i := range t {
			_, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				r1 = append(r1, i)
			} else {
				r2 = append(r2, i)
			}
		}
		if len(r1) > 0 {
			fmt.Print(strings.Join(r1, ","))
			if len(r2) > 0 {
				fmt.Print("|")
			}
		}
		if len(r2) > 0 {
			fmt.Print(strings.Join(r2, ","))
		}
		fmt.Println()
	}
}
