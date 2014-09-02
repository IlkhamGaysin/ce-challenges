package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var (
			n, v, vp int
			u        []bool
		)
		jolly, done := true, false
		fmt.Sscan(scanner.Text(), &n)
		if n == 1 {
			done = true
		} else {
			u = make([]bool, n-1)
		}
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &vp)
		for i := 1; i < n; i++ {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &v)
			if done {
				continue
			}
			x := abs(v - vp)
			if x >= n || x == 0 || u[x-1] {
				jolly, done = false, true
			} else {
				u[x-1], vp = true, v
			}
		}
		if jolly {
			fmt.Println("Jolly")
		} else {
			fmt.Println("Not jolly")
		}
	}
}
