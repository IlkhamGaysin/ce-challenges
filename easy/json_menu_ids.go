package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Item struct {
	Id    int
	Label string
}
type Menu struct {
	Header string
	Items  []Item
}
type Json struct {
	Menu Menu
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	reader := bufio.NewReader(data)
lines:
	for {
		var (
			s          []byte
			isPrefix   bool
			curr       Json
			sum_labels int
		)
		for {
			var sc []byte
			sc, isPrefix, err = reader.ReadLine()
			if err != nil {
				break lines
			}
			s = append(s, sc...)
			if !isPrefix {
				break
			}
		}
		if len(s) < 2 {
			continue
		}
		if err := json.Unmarshal(s, &curr); err != nil {
			print("Warning deserializing: ", fmt.Sprint(err))
		}
		for i := 0; i < len(curr.Menu.Items); i++ {
			var curr_label int
			fmt.Sscanf(curr.Menu.Items[i].Label, "Label %d", &curr_label)
			sum_labels += curr_label
		}
		fmt.Println(sum_labels)
	}
}
