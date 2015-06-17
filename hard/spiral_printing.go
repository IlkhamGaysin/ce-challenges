package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var row, col int
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var (
			r            []string
			i, j, tn, tw int
		)
		s := strings.Split(scanner.Text(), ";")
		fmt.Sscanf(s[0], "%d", &row)
		fmt.Sscanf(s[1], "%d", &col)
		t := strings.Fields(s[2])
		te, ts := col-1, row-1
		for {
			for j <= te {
				r = append(r, t[i*col+j])
				j++
			}
			j--
			i++
			tn++
			if i > ts {
				break
			}
			for i <= ts {
				r = append(r, t[i*col+j])
				i++
			}
			i--
			j--
			te--
			if j < tw {
				break
			}
			for j >= tw {
				r = append(r, t[i*col+j])
				j--
			}
			j++
			i--
			ts--
			if i < tn {
				break
			}
			for i >= tn {
				r = append(r, t[i*col+j])
				i--
			}
			i++
			j++
			tw++
			if j > te {
				break
			}
		}
		fmt.Println(strings.Join(r, " "))
	}
}
