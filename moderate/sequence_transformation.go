package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		s[0] = "^" + strings.Replace(strings.Replace(s[0], "0", "A+", -1), "1", "((A+)|(B+))", -1) + "$"
		seqRegex := regexp.MustCompile(s[0])
		if seqRegex.MatchString(s[1]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
