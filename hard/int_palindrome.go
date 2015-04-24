package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var t []int

func init() {
	t = make([]int, 10)
}

func pali(a int) bool {
	y := 0
	for c := a; c > 0; c /= 10 {
		t[y] = c % 10
		y++
	}
	for x := 0; ; x, y = x+1, y-1 {
		if t[x] != t[y-1] {
			return false
		}
		if y-x < 3 {
			return true
		}
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var l, r int
		fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		var n int
		for i := l; i <= r; i++ {
			prev := -1
			for j := i; j <= r; j++ {
				var p int
				if prev > -1 {
					p = prev
					if pali(j) {
						p++
					}
				} else {
					p = 0
					for k := i; k <= j; k++ {
						if pali(k) {
							p++
						}
					}
				}
				if p%2 == 0 {
					n++
				}
				prev = p
			}
		}
		fmt.Println(n)
	}
}
