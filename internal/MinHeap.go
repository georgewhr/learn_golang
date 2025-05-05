package internal

/*
MinHeap API:

insertVal()
removeVal()
getParent(), (childIndex -1) / 2
getLChild()
getRChild()
getMin() aee[0]
heapifyUp()
heapifyDown()







*/

type MinHeap struct {
	backingArr []int
	size       int
	cap        int
}

func MinHeapConstructor(cap int) *MinHeap {
	return &MinHeap{size: 0, cap: cap, backingArr: make([]int, cap)}
}

/*
Insert value to the array,
1.insert to the end of array
2.perform heapifyup
*/
func (this *MinHeap) InsertVal(val int) int {
	if this.size == this.cap {
		return -1
	}

	this.backingArr[this.size] = val
	this.heapifyUp()
	return 0

}

/*
Remove smallest Value
1. Assign min value to a temp variable
2. assign last element to arr[0]
3. perform heapifydown
*/
func (this *MinHeap) RemoveVal() int {
	n := len(this.backingArr)
	this.swap(0, n-1)
	valueToRemove := this.backingArr[n-1]
	this.backingArr = this.backingArr[:n-1]
	this.heapifyDown(0, n-2)
	return valueToRemove

}

/*
1. Swap the minIndex and last index,
2. remove last index
3. arr will be [:n-2]
4. heapfy down
1. compare left and right, pikc the smaller one
2. swap the samlelr one with current

*/

func (this *MinHeap) heapifyDown(currentIndex int, endIndex int) {
	leftChildIndex := 2*currentIndex + 1
	for leftChildIndex <= endIndex {
		rightChildIndex := 2*currentIndex + 2
		indexToSwap := leftChildIndex

		if this.backingArr[rightChildIndex] > this.backingArr[leftChildIndex] {
			indexToSwap = rightChildIndex
		}

		if this.backingArr[indexToSwap] < this.backingArr[currentIndex] {
			this.swap(indexToSwap, currentIndex)
			currentIndex = indexToSwap
			leftChildIndex = currentIndex*2 + 1
		}

	}

}

/*
Add value


*/

func (this *MinHeap) heapifyUp() {
	currentIndex := len(this.backingArr) - 1
	parentIndex := (currentIndex - 1) / 2
	for currentIndex > 0 && this.backingArr[parentIndex] < this.backingArr[currentIndex] {
		this.swap(currentIndex, parentIndex)
		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2
	}

}

func (this *MinHeap) swap(i int, j int) {

	temp := this.backingArr[i]
	this.backingArr[i] = this.backingArr[j]
	this.backingArr[j] = temp

}
