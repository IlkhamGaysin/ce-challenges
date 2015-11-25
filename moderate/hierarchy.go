package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Uint8s []uint8

func (s Uint8s) Len() int { return len(s) }
func (s Uint8s) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Uint8s) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func contains(a []uint8, b uint8) bool {
	for _, i := range a {
		if b == i {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var root uint8
		m := make(map[uint8][]uint8)
		parents, children := []uint8{}, []uint8{}
		s := strings.Split(scanner.Text(), " | ")
		for _, i := range s {
			m[i[0]] = append(m[i[0]], i[1])
			if !contains(parents, i[0]) {
				parents = append(parents, i[0])
			}
			if !contains(children, i[1]) {
				children = append(children, i[1])
			}
		}
		sort.Sort(Uint8s(parents))
		sort.Sort(Uint8s(children))
		for k := range m {
			sort.Sort(Uint8s(m[k]))
		}
		for _, i := range parents {
			if !contains(children, i) {
				root = i
				break
			}
		}
		r, c := string(root), root
		children = []uint8{}
		for len(m) > 0 {
			t := string(c)
			if len(m[c]) >= 1 {
				children = append(children, m[c]...)
				s = make([]string, len(m[c]))
				for ix, i := range m[c] {
					s[ix] = string(i)
				}
				t += " [" + strings.Join(s, ", ") + "]"
			}
			r = strings.Replace(r, string(c), t, -1)
			delete(m, c)
			for _, i := range children {
				if len(m[i]) > 0 {
					c = i
					break
				}
			}
		}
		fmt.Println(r)
	}
}
