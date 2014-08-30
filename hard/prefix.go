package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func pol(o byte, a, b float32) float32 {
	switch o {
	case '*':
		return a * b
	case '/':
		return a / b
	default: // '+'
		return a + b
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		n := len(s)
		var res float32
		fmt.Sscanf(s[n/2], "%f", &res)
		for i := 1; i <= n/2; i++ {
			var v float32
			fmt.Sscanf(s[n/2+i], "%f", &v)
			res = pol(s[n/2-i][0], res, v)
		}
		fmt.Println(int(res + 0.001))
	}
}
