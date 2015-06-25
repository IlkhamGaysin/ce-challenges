package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	emailRegex := regexp.MustCompile(`^((\"[0-9A-Za-z@.+-=]+\")|([0-9A-Za-z.+-=]+))@\w*(\w+\.)+\w{2,4}$`)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if emailRegex.MatchString(scanner.Text()) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
