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
	reader := bufio.NewReader(data)
	for {
		s, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		var n int
		fmt.Sscanf(string(s), "%d", &n)
		if n < 2 {
			fmt.Println(n)
		} else {
			b, c := []int64{1, 1, 0}, 1
			for n > c {
				c++
				b[0], b[2] = b[1]+b[2], b[1]
				b[1] = b[0]
			}
			fmt.Println(b[0])
		}
	}
}
