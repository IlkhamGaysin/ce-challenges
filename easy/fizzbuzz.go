package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	value string
	next  *node
}

func (a node) isEmpty() bool {
	return len(a.value) == 0
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var fizz, buzz, n int
		fmt.Sscanf(scanner.Text(), "%d %d %d", &fizz, &buzz, &n)

		tail := node{value: "FB"}
		list := &tail
		for a, b := fizz-1, buzz-1; a > 0 || b > 0; a, b = (a+fizz-1)%fizz, (b+buzz-1)%buzz {
			var value string
			if a == 0 {
				value = "F"
			} else if b == 0 {
				value = "B"
			}
			list = &node{value, list}
		}
		tail.next = list

		for i := 1; i <= n; i++ {
			if list.isEmpty() {
				fmt.Print(i)
			} else {
				fmt.Print(list.value)
			}
			list = list.next
			if i < n {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}
