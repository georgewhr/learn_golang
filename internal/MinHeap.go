package internal

/*
MinHeap API:

insertVal()
removeVal()
getParent(), childIndex/2
getLChild()
getRChild()
getMin()
heapifyUp()
heapifyDown()


parentIndex = 2n-1






*/

type MinHeap struct {
	backingArr []int
	size       int
	cap        int
}

func MinHeapConstructor(size int, cap int) *MinHeap {
	return &MinHeap{size: size, cap: cap, backingArr: make([]int, size)}
}

/*
Insert value to the array,
1.insert to the end of array
2.perform heapifyup
*/
func (this *MinHeap) InsertVal(val int) int {
	return -1

}

/*
Remove smallest Value
1. Assign min value to a temp variable
2. assign last element to arr[1]
3. perform heapifydown
*/
func (this *MinHeap) RemoveVal() int {
	return -1
}

/*
Remove min value, re-assign the end of index to beginging
*/

func (this *MinHeap) heapifyDown() {

}

/*
Add value


*/

func (this *MinHeap) heapifyUp() {

}
