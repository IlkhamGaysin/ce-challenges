package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func inCircle(a, b, c float64) bool {
	return a*a+b*b <= c*c
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		f := make([]float64, 5)
		fmt.Sscanf(scanner.Text(), "Center: (%f, %f); Radius: %f; Point: (%f, %f)",
			&f[0], &f[1], &f[2], &f[3], &f[4])
		if inCircle(f[0]-f[3], f[1]-f[4], f[2]) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
