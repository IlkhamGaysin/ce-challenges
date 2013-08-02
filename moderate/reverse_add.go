package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func rev(a int) (ret int) {
	for ; a > 0; a /= 10 {
		ret = 10*ret + a%10
	}
	return ret
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
		var a, b int
		fmt.Sscan(strings.TrimSpace(s), &a)
		for b < 100 {
			r := rev(a)
			if r == a {
				break
			}
			a, b = a+r, b+1
		}
		if b < 100 {
			fmt.Println(b, a)
		} else {
			fmt.Println("not found")
		}
	}
}
