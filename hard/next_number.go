package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func dsig(a int) (r int) {
	for a > 0 {
		if a%10 > 0 {
			r += 1 << uint(3*(a%10))
		}
		a /= 10
	}
	return r
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var d int
		fmt.Sscanf(scanner.Text(), "%d", &d)
		e := d + 9
		for ds := dsig(d); dsig(e) != ds; e += 9 {
		}
		fmt.Println(e)
	}
}
