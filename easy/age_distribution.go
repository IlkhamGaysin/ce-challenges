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
		var cat int
		for cat < len(age) && m >= age[cat] {
			cat++
		}
		fmt.Println(category[cat%len(age)])
	}
}
