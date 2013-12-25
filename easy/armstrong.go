package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewReader(data)
	var s string
	for {
		s, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		m, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatal(err)
		}
		var a, e, t, r int
		for a = m; a > 0; a /= 10 {
			e++
		}
		for a = m; a > 0; a /= 10 {
			r = a % 10
			t += powi(r, e)
		}
		if m == t {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
