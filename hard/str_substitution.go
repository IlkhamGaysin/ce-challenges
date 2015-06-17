package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type segment struct {
	s string
	r bool
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		t := strings.Split(s[1], ",")
		trans := []segment{segment{s[0], false}}
		for len(t) > 0 {
			for i := len(trans) - 1; i >= 0; i-- {
				if trans[i].r == false && strings.Contains(trans[i].s, t[0]) {
					u := strings.Split(trans[i].s, t[0])
					var v []segment
					for jx, j := range u {
						if len(j) > 0 {
							v = append(v, segment{j, false})
						}
						if jx != len(u)-1 {
							v = append(v, segment{t[1], true})
						}
					}
					if i > 0 {
						if i < len(trans)-1 {
							trans = append(trans[0:i], append(v, trans[i+1:len(trans)]...)...)
						} else {
							trans = append(trans[0:i], v...)
						}
					} else {
						if i < len(trans)-1 {
							trans = append(v, trans[i+1:len(trans)]...)
						} else {
							trans = v
						}
					}
				}
			}
			t = t[2:len(t)]
		}
		for _, i := range trans {
			fmt.Print(i.s)
		}
		fmt.Println()
	}
}
