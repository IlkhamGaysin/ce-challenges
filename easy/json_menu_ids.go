package main

import (
	"bufio"
	"encoding/json"
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
	reader := bufio.NewReader(data)
lines:
	for {
		var (
			s        []byte
			isPrefix bool
			curr     struct {
				Menu struct {
					Header string
					Items  []struct {
						Id    int
						Label string
					}
				}
			}
			sumLabels int
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
			var currLabel int
			fmt.Sscanf(curr.Menu.Items[i].Label, "Label %d", &currLabel)
			sumLabels += currLabel
		}
		fmt.Println(sumLabels)
	}
}
