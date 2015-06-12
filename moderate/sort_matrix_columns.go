package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type column []int
type Columns []column

func (slice Columns) Len() int { return len(slice) }
func (slice Columns) Less(i, j int) bool {
	var k int
	for k < len(slice[i])-1 && slice[i][k] == slice[j][k] {
		k++
	}
	return slice[i][k] < slice[j][k]
}
func (slice Columns) Swap(i, j int) { slice[i], slice[j] = slice[j], slice[i] }

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var i, j int
		s := strings.Split(scanner.Text(), " | ")
		n := strings.Count(s[0], " ") + 1
		m := make([]column, n)
		for i = 0; i < n; i++ {
			m[i] = make(column, n)
		}
		for i = 0; i < n; i++ {
			t := strings.Fields(s[i])
			for j = 0; j < n; j++ {
				fmt.Sscanf(t[j], "%d", &m[j][i])
			}
		}
		sort.Sort(Columns(m))
		r := make([]string, n)
		for i = 0; i < n; i++ {
			t := make([]string, n)
			for j = 0; j < n; j++ {
				t[j] = fmt.Sprint(m[j][i])
			}
			r[i] = strings.Join(t, " ")
		}
		fmt.Println(strings.Join(r, " | "))
	}
}
