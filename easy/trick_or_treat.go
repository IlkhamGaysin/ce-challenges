package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var v, z, w, h int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(),
			"Vampires: %d, Zombies: %d, Witches: %d, Houses: %d",
			&v, &z, &w, &h)
		fmt.Printf("%d\n", (v*3+z*4+w*5)*h/(v+z+w))
	}
}
