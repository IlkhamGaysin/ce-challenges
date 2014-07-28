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
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var num, count int
		sequ := strings.Fields(strings.TrimSpace(s))
		for _, i := range sequ {
			k, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			if k == num {
				count++
			} else {
				if count > 0 {
					fmt.Print(count, num, " ")
				}
				count, num = 1, k
			}
		}
		fmt.Println(count, num)
	}
}
