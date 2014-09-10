package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func lwc(f []string, c string) int {
	m := len(c) / 2
	if len(f) == 0 {
		return m
	}
	if len(c) == 0 {
		for ix, i := range f {
			f2 := make([]string, len(f))
			copy(f2, f)
			f2[ix], f2[len(f2)-1], f2 = f2[len(f2)-1], "", f2[:len(f2)-1]
			l := lwc(f2, i)
			if l > m {
				m = l
			}
		}
		return m
	}
	for ix, i := range f {
		if i[1] == c[0] {
			f2 := make([]string, len(f))
			copy(f2, f)
			f2[ix], f2[len(f2)-1], f2 = f2[len(f2)-1], "", f2[:len(f2)-1]
			l := lwc(f2, i+c)
			if l > m {
				m = l
			}
		}
		if i[0] == c[len(c)-1] {
			f2 := make([]string, len(f))
			copy(f2, f)
			f2[ix], f2[len(f2)-1], f2 = f2[len(f2)-1], "", f2[:len(f2)-1]
			l := lwc(f2, c+i)
			if l > m {
				m = l
			}
		}
	}
	return m
}

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
		t := strings.Split(strings.TrimSpace(s), ",")
		if len(t) < 2 {
			fmt.Println("None")
			continue
		}
		for ix, i := range t {
			t[ix] = string(i[0]) + string(i[len(i)-1])
		}
		l := lwc(t, "")
		if l == 1 {
			fmt.Println("None")
		} else {
			fmt.Println(l)
		}
	}
}
