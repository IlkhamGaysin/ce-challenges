package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type number struct {
	f float64
	s string
}
type Numbers []number

func (slice Numbers) Len() int           { return len(slice) }
func (slice Numbers) Less(i, j int) bool { return slice[i].f < slice[j].f }
func (slice Numbers) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		t := strings.Fields(scanner.Text())
		nums := []number{}
		for _, i := range t {
			var f float64
			fmt.Sscan(i, &f)
			nums = append(nums, number{f, i})
		}
		sort.Sort(Numbers(nums))
		for ix, i := range nums {
			if ix > 0 {
				fmt.Print(" ")
			}
			fmt.Print(i.s)
		}
		fmt.Println()
	}
}
