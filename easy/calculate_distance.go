package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		var x1, y1, x2, y2 int
		fmt.Sscanf(string(s), "(%d, %d) (%d, %d)", &x1, &y1, &x2, &y2)
		x, y := x1-x2, y1-y2
		fmt.Println(int(math.Sqrt(float64(x*x + y*y))))
	}
}
