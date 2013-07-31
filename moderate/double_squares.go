package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

func main() {
	sq := make(map[int]bool)
	for i := 0; i < 46341; i++ {
		sq[i*i] = true
	}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var n int
	fmt.Sscan(s, &n)
	for i := 0; i < n; i++ {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var x int
		fmt.Sscan(s, &x)
		num, l, bot, top := 0, make(map[int]bool), sqrt(x/2), sqrt(x)+1
		for j := bot; j < top; j++ {
			t := x - j*j
			if sq[t] && !l[t] {
				l[j*j], num = true, num+1
			}
		}
		fmt.Println(num)
	}
}
