package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const inf = math.MaxInt32

type vertex struct {
	id   uint
	dist uint
}

type edge struct {
	to   uint
	dist uint
}

type vertexHeap []vertex

func (h vertexHeap) Len() int           { return len(h) }
func (h vertexHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h vertexHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *vertexHeap) Push(x interface{}) {
	*h = append(*h, x.(vertex))
}
func (h *vertexHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h vertexHeap) FindID(a uint) int {
	for ix, i := range h {
		if i.id == a {
			return ix
		}
	}
	return -1
}

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var d uint = inf
		s := strings.Split(scanner.Text(), " | ")
		t := strings.Split(s[0], ", ")
		q := &vertexHeap{}
		heap.Init(q)
		neigh := make(map[uint][]edge)
		var a, b, v, x, y, z uint
		fmt.Sscanf(s[1], "%d %d", &a, &b)
		heap.Push(q, vertex{a, 0})
		v |= 1 << a
		if b != a {
			heap.Push(q, vertex{b, inf})
			v |= 1 << b
		}
		for _, i := range t {
			fmt.Sscanf(i, "%d %d %d", &x, &y, &z)
			neigh[x] = append(neigh[x], edge{y, z})
			neigh[y] = append(neigh[y], edge{x, z})
			if v&(1<<x) == 0 {
				heap.Push(q, vertex{x, inf})
				v |= 1 << x
			}
			if v&(1<<y) == 0 {
				heap.Push(q, vertex{y, inf})
				v |= 1 << y
			}
		}

		for q.Len() > 0 {
			u := heap.Pop(q).(vertex)
			if u.id == b {
				d = u.dist
				break
			}
			for _, i := range neigh[u.id] {
				dis := u.dist + i.dist
				ix := q.FindID(i.to)
				if ix >= 0 && dis < (*q)[ix].dist {
					(*q)[ix].dist = dis
					heap.Fix(q, ix)
				}
			}
		}

		if d == inf {
			fmt.Println("False")
		} else {
			fmt.Println(d)
		}
	}
}
