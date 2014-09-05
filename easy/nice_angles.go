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
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var a float64
		fmt.Sscanf(scanner.Text(), "%f", &a)
		fmt.Printf("%d.", int(a))
		a = (a - float64(int(a))) * 60
		fmt.Printf("%02d'", int(a))
		a = (a - float64(int(a))) * 60
		fmt.Printf("%02d\"\n", int(a))
	}
}
