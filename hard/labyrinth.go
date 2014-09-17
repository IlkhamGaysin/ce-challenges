package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type seg struct {
	pos  int
	from *seg
	to   []int
}

func kill(a seg, b int) {
	for ix, i := range a.to {
		if i == b {
			a.to[ix], a.to[len(a.to)-1], a.to = a.to[len(a.to)-1], 0, a.to[:len(a.to)-1]
			break
		}
	}
	if len(a.to) == 0 {
		kill(*a.from, a.pos)
	}
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var m map[int]bool
	var n int
	scanner := bufio.NewScanner(data)
	for i := 0; scanner.Scan(); i++ {
		if n == 0 {
			n = len(scanner.Text())
			m = make(map[int]bool, n*n)
		}
		for jx, j := range scanner.Text() {
			if j == ' ' {
				m[i*n+jx] = true
			}
		}
	}
	north, south := []seg{}, []seg{}
	for i := 0; i < n; i++ {
		if m[i] && m[i+n] {
			start := seg{i, nil, []int{i + n}}
			north = append(north, seg{i + n, &start, []int{}})
		}
		if m[(n-1)*n+i] && m[(n-2)*n+i] {
			start := seg{(n-1)*n + i, nil, []int{(n-2)*n + i}}
			south = append(south, seg{(n-2)*n + i, &start, []int{}})
		}
	}
	path := []int{}
	for {
		for i, nr := 0, len(north); i < nr; i++ {
			curr := north[0]
			copy(north, north[1:])
			north = north[:len(north)-1]
			count := 0
			checks := []int{curr.pos + n, curr.pos - 1, curr.pos + 1, curr.pos - n}
			for _, check := range checks {
				if check != curr.from.pos {
					if _, open := m[check]; open {
						var f bool
						for _, j := range north {
							if check == j.pos {
								f = true
								break
							}
						}
						if f {
							break
						}
						for _, j := range south {
							if check == j.pos {
								for ; curr.pos > n; curr = *curr.from {
									path = append([]int{curr.pos}, path...)
								}
								path = append([]int{curr.pos}, path...)
								for curr = j; curr.pos < (n-1)*n; curr = *curr.from {
									path = append(path, curr.pos)
								}
								path = append(path, curr.pos)
								goto out
							}
						}
						next := seg{check, &curr, []int{}}
						north = append(north, next)
						curr.to = append(curr.to, check)
						count++
					}
				}
			}
			if count == 0 {
				kill(*curr.from, curr.pos)
			}
		}

		for i, sr := 0, len(south); i < sr; i++ {
			curr := south[0]
			copy(south, south[1:])
			south = south[:len(south)-1]
			count := 0
			checks := []int{curr.pos + n, curr.pos - 1, curr.pos + 1, curr.pos - n}
			for _, check := range checks {
				if check != curr.from.pos {
					if _, open := m[check]; open {
						var f bool
						for _, j := range south {
							if check == j.pos {
								f = true
								break
							}
						}
						if f {
							break
						}
						for _, j := range north {
							if check == j.pos {
								for ; curr.pos < (n-1)*n; curr = *curr.from {
									path = append(path, curr.pos)
								}
								path = append(path, curr.pos)
								for curr = j; curr.pos > n; curr = *curr.from {
									path = append([]int{curr.pos}, path...)
								}
								path = append([]int{curr.pos}, path...)
								goto out
							}
						}
						next := seg{check, &curr, []int{}}
						south = append(south, next)
						curr.to = append(curr.to, check)
						count++
					}
				}
			}
			if count == 0 {
				kill(*curr.from, curr.pos)
			}
		}
	}
out:
	sort.Ints(path)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i*n+j] {
				for path[count] < i*n+j {
					count++
				}
				if path[count] == i*n+j {
					fmt.Print("+")
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
