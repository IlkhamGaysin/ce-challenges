package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func happy(a uint) (ret uint) {
	for a > 0 {
		b := a % 10
		ret += b * b
		a /= 10
	}
	return ret
}

func main() {
	var a uint
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &a)
		b := []uint{a}
		for i := 0; a > 1 && i < 127; i++ {
			a = happy(a)
			for _, j := range b {
				if j == a {
					a = 0
					break
				}
			}
			b = append(b, a)
		}
		if a == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
