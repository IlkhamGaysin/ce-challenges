package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const k = ` !"#$%&'()*+,-./0123456789:<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ";")
		for i := 0; i < len(s[1]); i++ {
			p := strings.IndexRune(k, rune(s[1][i])) - int(s[0][i%len(s[0])]-'0')
			fmt.Printf("%c", k[(len(k)+p)%len(k)])
		}
		fmt.Println()
	}
}
