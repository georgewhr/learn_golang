/*
Key note of a HashMap
1.Has a backing array with a initialized size
2.Handle collision, each index is the hash code, then append the list with key and value
3.When to extend the size? define a load factor, resize hashtable will happen if 1), collision to much 2) bucket is near factor

Method:
Put(k, v), null or previous value
Get(k), null or non-nil value
hash()
resize()

concurrent hashmap

e.g, size of 4 backing array, put 4/n sections, n = 2,  each section has a lock,
0,1  2, 3, construct a map where key is the actual hashkey, value is the sections.
hashkey = hash() % len(array), hashkey % (len(arr)/n)




*/

package internal

import "sync"

const TABLE_INIT_CAP = 50
const THREAD_SHOLD = 0.7

type HashNode struct {
	key  int
	val  int
	next *HashNode
}

type HashMap struct {
	BackingArr  []*HashNode
	tableSize   int
	tableCap    int
	sectionLock []sync.Mutex
}

func InitHashMap() *HashMap {
	return &HashMap{tableSize: 0,
		tableCap: TABLE_INIT_CAP, BackingArr: make([]*HashNode, TABLE_INIT_CAP),
	}
}

func (h *HashMap) Put(key int, value int) int {
	index := h.getIndex(key)

	if float64(h.tableSize)/float64(h.tableCap) > 0.6 {
		h.resize()
	}

	if h.BackingArr[index] == nil {
		h.BackingArr[index] = &HashNode{key: key, val: value}
	} else {
		headNode := h.BackingArr[index]
		for headNode.next != nil && headNode.key != key {
			headNode = headNode.next
		}

		if headNode.key == key {
			previousVal := headNode.key
			headNode.val = value
			return previousVal
		} else {
			headNode.next = &HashNode{key: key, val: value}
			return 0

		}
	}
	return -1

}

func (h *HashMap) Get(key int, value int) int {
	index := h.getIndex(key)

	if h.BackingArr[index] == nil {
		return -1
	}

	temp := h.BackingArr[index]

	for temp != nil && temp.key != key {
		temp = temp.next
	}

	if temp == nil {
		return -1
	} else {
		return temp.val
	}

	return 0

}

func (h *HashMap) getIndex(key int) (index int) {
	return int(hash(string(key))) % h.tableCap
}

func hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return
}

/*
if size == cap
  allocate array with 2*cap

copy from orignal array to new array
update array pointer


*/

func (h *HashMap) resize() {
	newArr := make([]*HashNode, h.tableSize, 2*h.tableCap)
	copy(newArr, h.BackingArr[:])
	h.BackingArr = newArr

}

/*
3 1/2 = 0
3 3, 2/2=1
3 3 3, 3, 3/2 = 1


3 3 3 3, 4, 4/2= 2
3 3 3 3 3, 5, 5/2 = 2
3 3 3 3 3 3, 6 /2  = 3
3 3 3 3 3 3 3, 7 /2  = 3

*/
