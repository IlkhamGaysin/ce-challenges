package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value int
	next  *node
}

func main() {
	var n, m int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var t []string
		fmt.Sscanf(scanner.Text(), "%d,%d", &n, &m)

		tail := node{value: n - 1}
		list := &tail
		for i := n - 2; i >= 0; i-- {
			list = &node{i, list}
		}
		tail.next, list = list, &tail

		for i := 0; i < n; i++ {
			for j := 0; j < m-1; j++ {
				list = list.next
			}
			t = append(t, strconv.Itoa(list.next.value))
			list.next = list.next.next
		}
		fmt.Println(strings.Join(t, " "))
	}
}
