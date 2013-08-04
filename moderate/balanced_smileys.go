package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func isBalanced(s string, c int) bool {
	for {
		if c < 0 {
			return false
		} else if s == "" {
			return c == 0
		}
		m1, err1 := regexp.Match(`[a-z ]`, []byte{s[0]})
		m2, err2 := regexp.Match(`[a-z :]`, []byte{s[len(s)-1]})
		switch {
		case m1:
			s = s[1:]
		case m2:
			s = s[:len(s)-1]
		case s[0] == '(':
			c, s = c+1, s[1:]
		case s[0] == ')':
			c, s = c-1, s[1:]
		case s[0] == ':':
			if s[1] == '(' {
				return isBalanced(s[2:], c) || isBalanced(s[2:], c+1)
			} else if s[1] == ')' {
				return isBalanced(s[2:], c) || isBalanced(s[2:], c-1)
			} else {
				s = s[1:]
			}
		case err1 != nil || err2 != nil:
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
