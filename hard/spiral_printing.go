package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		s := strings.Split(scanner.Text(), ";")
		var row, col int
		fmt.Sscanf(s[0], "%d", &row)
		fmt.Sscanf(s[1], "%d", &col)
		t := strings.Fields(s[2])
		i, j := 0, 0
		tn, te, ts, tw := 0, col-1, row-1, 0
		r := []string{}
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
