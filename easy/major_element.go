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
		var maxnum, maxcount int
		num := make([]int, 101)
		sequ := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		for ix, i := range sequ {
			k, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			num[k]++
			if num[k] > maxcount {
				maxnum, maxcount = k, num[k]
			}
			if len(sequ)/2+1 > maxcount+len(sequ)-ix-1 {
				fmt.Println("None")
				break
			} else if maxcount >= len(sequ)/2+1 {
				fmt.Println(maxnum)
				break
			}
		}
	}
}
