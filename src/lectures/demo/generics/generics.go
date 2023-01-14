package main

import "fmt"

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[P comparable, V any] struct {
	items      map[P][]V
	priorities []P
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{
		items:      make(map[P][]V),
		priorities: priorities,
	}
}

func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[P, V]) Next() (V, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		items := pq.items[priority]
		if len(items) > 0 {
			next := items[0]
			pq.items[priority] = items[1:]
			return next, true
		}
	}
	var empty V
	return empty, false
}

func (pq *PriorityQueue[P, V]) Empty() bool {
	for i := 0; i < len(pq.priorities); i++ {
		if len(pq.items[pq.priorities[i]]) > 0 {
			return false
		}
	}
	return true
}

func main() {
	priorities := []int{High, Medium, Low}
	pq := NewPriorityQueue[int, string](priorities)
	pq.Add(High, "high text 1")
	pq.Add(High, "high text 2")
	pq.Add(Medium, "medium text")
	pq.Add(Low, "low text")
	for !pq.Empty() {
		str, _ := pq.Next()
		fmt.Printf("pq top string is '%s'\n", str)
	}
}
