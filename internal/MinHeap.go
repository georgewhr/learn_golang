package internal

/*
MinHeap API:

insert()
getParent()
getLChild()
getRChild()
getMin()
peek()
poll()
minheapify()





*/


type MinHeap struct {
	backingArr []int
	size       int
	cap        int
}

func MinHeapConstructor(size int, cap int) *MinHeap {
	return &MinHeap{size: size, cap: cap, backingArr: make([]int, size)}
}


func (this *MinHeap) 
