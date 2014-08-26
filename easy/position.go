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
		var n, p, q uint
		fmt.Sscanf(string(s), "%d,%d,%d", &n, &p, &q)
		if ((n & (1 << (p - 1))) == 0) == ((n & (1 << (q - 1))) == 0) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
