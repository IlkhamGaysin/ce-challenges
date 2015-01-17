package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	category := []string{"This program is for humans",
		"Still in Mama's arms",
		"Preschool Maniac",
		"Elementary school",
		"Middle school",
		"High school",
		"College",
		"Working for the man",
		"The Golden Years"}
	age := []int{0, 3, 5, 12, 15, 19, 23, 66, 101}

	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var cat, m int
		fmt.Sscanf(scanner.Text(), "%d", &m)
		for cat < len(age) && m >= age[cat] {
			cat++
		}
		fmt.Println(category[cat%len(age)])
	}
}
