package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	units := []string{"PENNY", "NICKEL", "DIME", "QUARTER", "HALF DOLLAR",
		"ONE", "TWO", "FIVE", "TEN", "TWENTY", "FIFTY", "ONE HUNDRED"}
	value := []int{1, 5, 10, 25, 50, 100,
		200, 500, 1000, 2000, 5000, 10000}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		var p, c, v int
		if strings.Contains(s[0], ".") {
			fmt.Sscanf(s[0], "%d.%d", &p, &v)
		} else {
			fmt.Sscanf(s[0], "%d", &p)
		}
		p, v = p*100+v, 0
		if strings.Contains(s[1], ".") {
			fmt.Sscanf(s[1], "%d.%d", &c, &v)
		} else {
			fmt.Sscanf(s[1], "%d", &c)
		}
		c = c*100 + v
		if p > c {
			fmt.Println("ERROR")
		} else if p == c {
			fmt.Println("ZERO")
		} else {
			var res []string
			change := make([]int, 12)
			for i := 11; i >= 0; i-- {
				for c-p >= value[i] {
					change[i]++
					c -= value[i]
				}
			}
			for i := 11; i >= 0; i-- {
				for change[i] > 0 {
					res = append(res, units[i])
					change[i]--
				}
			}
			fmt.Println(strings.Join(res, ","))
		}
	}
}
