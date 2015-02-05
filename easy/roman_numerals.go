package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	ronum, rostr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1},
		[]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var num int
		fmt.Sscan(scanner.Text(), &num)
		for i := 0; num > 0; {
			if num >= ronum[i] {
				fmt.Print(rostr[i])
				num -= ronum[i]
			} else {
				i++
			}
		}
		fmt.Println()
	}
}
