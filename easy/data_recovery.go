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
		t := strings.Split(scanner.Text(), ";")
		if len(t) < 2 {
			continue
		}
		words, sequ := strings.Fields(t[0]), strings.Fields(t[1])
		reco := make([]string, len(words))
		for ix, i := range sequ {
			k, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			reco[k-1] = words[ix]
		}
		for ix := 0; ix < len(reco); ix++ {
			if reco[ix] == "" {
				reco[ix] = words[len(words)-1]
				break
			}
		}
		fmt.Println(strings.Join(reco, " "))
	}
}
