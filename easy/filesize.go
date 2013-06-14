package main

import (
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
	stat, err := data.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stat.Size())
}
