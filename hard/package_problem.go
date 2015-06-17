package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type item struct {
	id     int64
	value  int64
	weight int64
}
type Items []item

func (slice Items) Len() int { return len(slice) }
func (slice Items) Less(i, j int) bool {
	return (float64(slice[i].value) / float64(slice[i].weight)) > (float64(slice[j].value) / float64(slice[j].weight))
}
func (slice Items) Swap(i, j int) { slice[i], slice[j] = slice[j], slice[i] }

type Best []int64

func (slice Best) Len() int           { return len(slice) }
func (slice Best) Less(i, j int) bool { return slice[i] < slice[j] }
func (slice Best) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

type task struct {
	level int64
	bound int64
	value int64
	room  int64
	items []int64
}

func boundary(room, value int64, items []item, itl []int64) (t, b int64, tl []int64) {
	total_room, f := room, false
	t, b = value, value
	for _, i := range items {
		if i.weight > total_room {
			continue
		} else if i.weight > room {
			if !f {
				b += int64(float64(i.value*room) / float64(i.weight))
				f = true
			}
		} else {
			room, t, tl = room-i.weight, t+i.value, append(tl, i.id)
			if !f {
				b += i.value
			}
			if room == 0 {
				break
			}
		}
	}
	return t, b, tl
}

func main() {
	var (
		ks, s0, s2 int64
		s1         float64
	)
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if len(scanner.Text()) < 3 {
			continue
		}
		ts := strings.Split(scanner.Text(), "(")
		n := len(ts) - 1
		fmt.Sscanf(ts[0], "%d : ", &ks)
		var (
			items      []item
			min_weight int64 = math.MaxInt64
		)
		n0 := n
		for i := 1; i <= n0; i++ {
			fmt.Sscanf(ts[i], "%d,%f,$%d)", &s0, &s1, &s2)

			if s1 > float64(ks) {
				n--
			} else {
				items = append(items, item{s0, s2, int64(s1 * 100)})
			}
			if s2 < min_weight {
				min_weight = s2
			}
		}
		if n == 0 {
			fmt.Println("-")
			continue
		}
		sort.Sort(Items(items))
		ks *= 100

		t, bound, tl := boundary(ks, int64(0), items[1:], []int64{})
		best, bestl := t, make([]int64, len(tl))
		todo := []task{task{bound: bound, room: ks}}
		copy(bestl, tl)

		t, bound, tl = boundary(ks, int64(0), items, []int64{})
		if t > best {
			best, bestl = t, make([]int64, len(tl))
			copy(bestl, tl)
		}
		if items[0].weight <= ks {
			todo = append(todo, task{int64(0), bound, items[0].value, ks - items[0].weight, []int64{items[0].id}})
		}

		for len(todo) > 0 {
			curr := todo[len(todo)-1]
			todo = todo[:len(todo)-1]
			if curr.bound <= best {
				continue
			}

			if curr.level == int64(n-2) {
				if curr.room >= items[n-1].weight {
					t = items[n-1].value + curr.value
					if t > best {
						best, bestl = t, make([]int64, len(curr.items)+1)
						copy(bestl, append(curr.items, items[n-1].id))
					}
				}
				if curr.value > best {
					best, bestl = curr.value, make([]int64, len(curr.items))
					copy(bestl, curr.items)
				}
			} else {
				t, bound, tl = boundary(curr.room, curr.value, items[curr.level+1:], curr.items)
				if t > best {
					tl = append(curr.items, tl...)
					best, bestl = t, make([]int64, len(tl))
					copy(bestl, tl)
				}
				if bound > best && curr.room > min_weight {
					todo = append(todo, task{curr.level + 1, bound, curr.value, curr.room, curr.items})
				}

				if curr.room >= items[curr.level+1].weight+min_weight {
					todo = append(todo, task{curr.level + 1, curr.bound, curr.value + items[curr.level+1].value, curr.room - items[curr.level+1].weight, append(curr.items, items[curr.level+1].id)})
				}
			}
		}
		var st []string
		sort.Sort(Best(bestl))
		for i := int64(1); i <= int64(n0); i++ {
			if len(bestl) > 0 && i == bestl[0] {
				st = append(st, fmt.Sprint(i))
				bestl = bestl[1:]
			}
		}
		fmt.Println(strings.Join(st, ","))
	}
}
