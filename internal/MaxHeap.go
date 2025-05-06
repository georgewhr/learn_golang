package internal

/*
MaxHeap API:

insertVal()
removeVal()
getParent(), (childIndex -1) / 2
getLChild()
getRChild()
getMin() aee[0]
heapifyUp()
heapifyDown()




*/

type MaxHeap struct {
	backingArr []int
	size       int
	cap        int
}

func MaxHeapConstructor(cap int) *MaxHeap {
	return &MaxHeap{size: 0, cap: cap, backingArr: make([]int, 0)}
}

/*
Insert value to the array,
1.insert to the end of array
2.perform heapifyup
*/
func (this *MaxHeap) InsertVal(val int) int {

	this.backingArr = append(this.backingArr, val)
	this.heapifyUp()
	this.size++
	return 0

}

/*
Remove smallest Value
1. Assign min value to a temp variable
2. assign last element to arr[0]
3. perform heapifydown
*/
func (this *MaxHeap) RemoveVal() int {
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

func (this *MaxHeap) heapifyDown(currentIndex int, endIndex int) {
	leftChildIndex := 2*currentIndex + 1
	for leftChildIndex <= endIndex {
		rightChildIndex := 2*currentIndex + 2
		if rightChildIndex > endIndex {
			rightChildIndex = -1
		}
		indexToSwap := leftChildIndex

		if rightChildIndex != -1 && this.backingArr[rightChildIndex] > this.backingArr[leftChildIndex] {
			indexToSwap = rightChildIndex
		}

		if this.backingArr[indexToSwap] > this.backingArr[currentIndex] {
			this.swap(indexToSwap, currentIndex)
			currentIndex = indexToSwap
			leftChildIndex = currentIndex*2 + 1
		} else {
			return
		}

	}

}

/*
Add value


*/

func (this *MaxHeap) heapifyUp() {
	currentIndex := len(this.backingArr) - 1
	parentIndex := (currentIndex - 1) / 2
	for currentIndex > 0 && this.backingArr[parentIndex] < this.backingArr[currentIndex] {
		this.swap(currentIndex, parentIndex)
		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2
	}

}

func (this *MaxHeap) swap(i int, j int) {

	temp := this.backingArr[i]
	this.backingArr[i] = this.backingArr[j]
	this.backingArr[j] = temp

}
