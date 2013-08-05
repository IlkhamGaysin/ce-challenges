package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isBalanced(s string, c int) bool {
	for {
		if c < 0 {
			return false
		} else if s == "" {
			return c == 0
		}
		first, last := s[0], s[len(s)-1]
		switch {
		case (first >= 'a' && first <= 'z') || first == ' ':
			s = s[1:]
		case (last >= 'a' && last <= 'z') || last == ' ' || last == ':':
			s = s[:len(s)-1]
		case first == '(':
			c, s = c+1, s[1:]
		case first == ')':
			c, s = c-1, s[1:]
		case first == ':':
			if s[1] == '(' {
				return isBalanced(s[2:], c) || isBalanced(s[2:], c+1)
			} else if s[1] == ')' {
				return isBalanced(s[2:], c) || isBalanced(s[2:], c-1)
			} else {
				s = s[1:]
			}
		default:
			return false
		}
	}
	return false
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
		if isBalanced(strings.TrimSpace(s), 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
