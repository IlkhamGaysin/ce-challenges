package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func powi(a, b int) (ret int) {
	ret = 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var m, e, t int
		fmt.Sscanf(scanner.Text(), "%d", &m)
		for a := m; a > 0; a /= 10 {
			e++
		}
		for a := m; a > 0; a /= 10 {
			t += powi(a%10, e)
		}
		if m == t {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
