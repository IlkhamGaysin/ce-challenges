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
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanBytes)
	valid := true
	stack := []byte{}
	var c uint8
	for scanner.Scan() {
		c = scanner.Text()[0]
		if !valid {
			if c == '\n' {
				valid = true
			}
			continue
		}
		switch c {
		case '\n':
			if len(stack) > 0 {
				fmt.Println("False")
				valid = true
			} else {
				fmt.Println("True")
			}
		case ')':
			if len(stack) > 0 && stack[0] == '(' {
				stack = stack[1:len(stack)]
			} else {
				valid = false
			}
		case ']':
			if len(stack) > 0 && stack[0] == '[' {
				stack = stack[1:len(stack)]
			} else {
				valid = false
			}
		case '}':
			if len(stack) > 0 && stack[0] == '{' {
				stack = stack[1:len(stack)]
			} else {
				valid = false
			}
		default:
			stack = append([]byte{c}, stack...)
		}
		if !valid {
			fmt.Println("False")
			stack = []byte{}
		}
	}
	if c != '\n' {
		if len(stack) > 0 {
			fmt.Println("False")
		} else {
			fmt.Println("True")
		}
	}
}
