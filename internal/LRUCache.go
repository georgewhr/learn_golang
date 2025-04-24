package internal

/*
Implement a LRU cache
Requirements:
LRU cache, age out the least used item
init(capacity)
get, O(1)
put, O(1)


Put(val)
1. Check if size == cap, if yes, remove last item from the queue, then remove from map
2. Add new Item in the map, then insert into the queue


Get(key)
1. check if key exist
2. get from map return the list node
3. remove the list node, and then insert again.

*/

type LRUCache struct {
	size       int
	cap        int
	backingMap map[int]*ListNode
	double     *ListNode
	doubleTail *ListNode
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		cap:        capacity,
		size:       0,
		backingMap: make(map[int]*ListNode),
	}
}

/*
need to be O(1)

if key present {
	get node from map
	remove node from list
	add new node with new value again
} else {
	if cap == size {
		evict least recent item, which is the end of list node
	}
	add key
	add node to the front
}
*/

func (this *LRUCache) put(key int, value int) int {
	if _, ok := this.backingMap[key]; ok {
		oldNode := this.backingMap[key]
		this.removeNode(oldNode)
		this.backingMap[key] = this.addNode(value)
	} else {
		if this.cap == this.size {
			//remove end of node
		}
		this.backingMap[key] = this.addNode(value)

	}
	return 1
}

/*
need to be O(1)
if key present {
	get value
	remove node
	add node to front

} else {
	return -1

}
*/

func (this *LRUCache) get(key int) int {
	if val, ok := this.backingMap[key]; ok {
		this.removeNode(val)
		this.addNode(val.Val)
		return val.Val
	} else {
		return -1
	}
}

func (this *LRUCache) addNode(val int) *ListNode {
	newNode := &ListNode{Val: val, Next: nil}
	if this.double == nil {
		this.double = newNode
		this.doubleTail = this.double
	} else {
		newNode.Next = this.double
		this.double.Previous = newNode
		newNode = this.double
	}
	return newNode
}

func (this *LRUCache) removeNode(node *ListNode) {
	node.Previous.Next = node.Next
}

func (this *LRUCache) removeTail() {
	this.doubleTail.Previous.Next = nil
	this.doubleTail = this.doubleTail.Previous
}
