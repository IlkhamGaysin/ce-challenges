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
type numbers []number

func (slice numbers) Len() int           { return len(slice) }
func (slice numbers) Less(i, j int) bool { return slice[i].f < slice[j].f }
func (slice numbers) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func main() {
	var f float64
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var nums []number
		t := strings.Fields(scanner.Text())
		for _, i := range t {
			fmt.Sscan(i, &f)
			nums = append(nums, number{f, i})
		}
		sort.Sort(numbers(nums))
		for ix, i := range nums {
			if ix > 0 {
				fmt.Print(" ")
			}
			fmt.Print(i.s)
		}
		fmt.Println()
	}
}
