package internal

import "fmt"

// MaxHeap implements heap.Interface and holds Items
type MinHeapListNode []*ListNode

func (h MinHeapListNode) Len() int { return len(h) }
func (h MinHeapListNode) Less(i, j int) bool {
	fmt.Print("asd")
	return h[i].Val < h[j].Val
}                                       // minHeap
func (h MinHeapListNode) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeapListNode) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeapListNode) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
