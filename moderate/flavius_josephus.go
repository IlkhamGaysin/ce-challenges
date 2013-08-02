package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
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
		var n, m int
		fmt.Sscanf(s, "%d,%d", &n, &m)

		tail := Node{n - 1, nil}
		list := &tail
		for i := n - 2; i >= 0; i-- {
			list = &Node{i, list}
		}
		tail.next, list = list, &tail

		t := []string{}
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
