package internal

type TaskSchedule struct {
	ID       int
	Priority int
}

// MaxHeap implements heap.Interface and holds Items
type TaskScheduleMaxHeap []TaskSchedule

func (h TaskScheduleMaxHeap) Len() int           { return len(h) }
func (h TaskScheduleMaxHeap) Less(i, j int) bool { return h[i].Priority > h[j].Priority } // MaxHeap
func (h TaskScheduleMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskScheduleMaxHeap) Push(x any) {
	*h = append(*h, x.(TaskSchedule))
}

func (h *TaskScheduleMaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
