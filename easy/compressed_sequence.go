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
		var num, count, k int
		sequ := strings.Fields(strings.TrimSpace(scanner.Text()))
		for _, i := range sequ {
			fmt.Sscanf(i, "%d", &k)
			if k == num {
				count++
				continue
			}
			if count > 0 {
				fmt.Print(count, num, " ")
			}
			count, num = 1, k
		}
		fmt.Println(count, num)
	}
}
