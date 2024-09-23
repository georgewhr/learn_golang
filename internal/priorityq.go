package internal

import (
	"container/heap"
)

type Item struct {
	Value    int
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (piq PriorityQueue) Less(i, j int) bool {

	return piq[i].Priority < piq[j].Priority
}
func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
	piq[i].Index = i
	piq[j].Index = j
}
func (piq *PriorityQueue) Push(x interface{}) {
	n := len(*piq)
	item := x.(*Item)
	item.Index = n
	*piq = append(*piq, item)
}
func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*piq = old[0 : n-1]
	return item
}
func (piq *PriorityQueue) Update(item *Item, Value int, Priority int) {
	item.Value = Value
	item.Priority = Priority
	heap.Fix(piq, item.Index)
}
