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
		var (
			num              [101]int
			maxnum, maxcount int
		)
		sequ := strings.Split(strings.TrimSpace(s), ",")
		for ix, i := range sequ {
			k, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			num[k]++
			if num[k] > maxcount {
				maxnum = k
				maxcount = num[k]
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
