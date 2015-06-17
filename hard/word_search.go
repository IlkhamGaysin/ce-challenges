package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isContained(pos []int, bee []string, st0 string) bool {
	if st0 == "" {
		return true
	}
	var neigh [][]int
	if pos[0] > 0 && bee[pos[0]-1][pos[1]] == st0[0] {
		neigh = append(neigh, []int{pos[0] - 1, pos[1]})
	}
	if pos[0] < 2 && bee[pos[0]+1][pos[1]] == st0[0] {
		neigh = append(neigh, []int{pos[0] + 1, pos[1]})
	}
	if pos[1] > 0 && bee[pos[0]][pos[1]-1] == st0[0] {
		neigh = append(neigh, []int{pos[0], pos[1] - 1})
	}
	if pos[1] < 3 && bee[pos[0]][pos[1]+1] == st0[0] {
		neigh = append(neigh, []int{pos[0], pos[1] + 1})
	}
	if len(neigh) == 0 {
		return false
	}
	for _, nei := range neigh {
		be0 := make([]string, len(bee))
		copy(be0, bee)
		be0[pos[0]] = bee[pos[0]][:pos[1]] + " " + bee[pos[0]][pos[1]+1:]
		if isContained(nei, be0, st0[1:]) {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	board := []string{"ABCE", "SFCS", "ADEE"}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := scanner.Text()

		var (
			start [][]int
			f     bool
		)
		for ix, i := range board {
			for jx, j := range i {
				if rune(s[0]) == j {
					start = append(start, []int{ix, jx})
				}
			}
		}

		for _, i := range start {
			if isContained(i, board, s[1:]) {
				f = true
				break
			}
		}
		if f {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
