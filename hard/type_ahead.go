package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const text = `
Mary had a little lamb its fleece was white as snow;
And everywhere that Mary went, the lamb was sure to go. 
It followed her to school one day, which was against the rule;
It made the children laugh and play, to see a lamb at school.
And so the teacher turned it out, but still it lingered near,
And waited patiently about till Mary did appear.
"Why does the lamb love Mary so?" the eager children cry;
"Why, Mary loves the lamb, you know" the teacher did reply."
`

var freq map[string]int

type ngrams [][]string

func (s ngrams) Len() int { return len(s) }
func (s ngrams) Less(i, j int) bool {
	a, b := strings.Join(s[i], ","), strings.Join(s[j], ",")
	if freq[a] == freq[b] {
		return a < b
	}
	return freq[a] > freq[b]
}
func (s ngrams) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sliceq(a, b []string) bool {
	for ix, i := range a {
		if i != b[ix] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	words := strings.Fields(text)
	for ix, i := range words {
		words[ix] = strings.Trim(i, ";,.?\"")
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var tnl, r [][]string
		t := strings.Split(scanner.Text(), ",")
		x := strings.Fields(strings.TrimSpace(t[1]))
		fmt.Sscanf(t[0], "%d", &n)

		for i := 0; i <= len(words)-n; i++ {
			temp := make([]string, n)
			for j := 0; j < n; j++ {
				temp[j] = words[i+j]
			}
			tnl = append(tnl, temp)
		}

		for _, i := range tnl {
			f := true
			for j := 0; j < len(x); j++ {
				if i[j] != x[j] {
					f = false
					break
				}
			}
			if f {
				r = append(r, i)
			}
		}

		freq = make(map[string]int)
		for _, i := range r {
			rs := strings.Join(i, ",")
			if _, f := freq[rs]; f {
				freq[rs]++
			} else {
				freq[rs] = 1
			}
		}

		sort.Sort(ngrams(r))
		for ix, i := range r {
			if ix == 0 || !sliceq(i, r[ix-1]) {
				if ix > 0 {
					fmt.Print(";")
				}
				fmt.Print(strings.Join(i[len(x):], " "))
				fmt.Printf(",%.3f", float32(freq[strings.Join(i, ",")])/float32(len(r)))
			}
		}
		fmt.Println()
	}
}
