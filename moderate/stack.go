package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func stackpush(stack []int, a int) []int {
	return append(stack, a)

}

func stackpop(stack []int) ([]int, int) {
	var a int
	if len(stack) > 0 {
		a, stack = stack[len(stack)-1], stack[:len(stack)-1]
	}
	return stack, a
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
		t, stack := strings.Fields(strings.TrimSpace(s)), []int{}
		for _, i := range t {
			var u int
			fmt.Sscan(i, &u)
			stack = stackpush(stack, u)
		}

		if len(stack) > 0 {
			var item int
			stack, item = stackpop(stack)
			fmt.Print(item)
			stack, _ = stackpop(stack)
		}
		for len(stack) > 0 {
			var item int
			stack, item = stackpop(stack)
			fmt.Print(" ", item)
			stack, _ = stackpop(stack)
		}
		fmt.Println()
	}
}
