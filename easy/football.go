package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type team struct {
	id int
	cs string
}

type Teams []team

func (slice Teams) Len() int           { return len(slice) }
func (slice Teams) Less(i, j int) bool { return slice[i].id < slice[j].id }
func (slice Teams) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func main() {
	var a int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var (
			r []team
			q []string
		)
		m := make(map[int][]string)
		s := strings.Split(scanner.Text(), "|")
		for ix, i := range s {
			t := strings.Fields(i)
			for _, j := range t {
				fmt.Sscanf(j, "%d", &a)
				m[a] = append(m[a], fmt.Sprint(ix+1))
			}
		}
		for k, v := range m {
			r = append(r, team{k, strings.Join(v, ",") + ";"})
		}
		sort.Sort(Teams(r))
		for _, i := range r {
			q = append(q, fmt.Sprintf("%d:%s", i.id, i.cs))
		}
		fmt.Println(strings.Join(q, " "))
	}
}
