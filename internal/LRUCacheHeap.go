package internal

/*
LRU cache, age out the least used item
init(capacity)
get, O(1)
put, O(1)


Put(val)
1. make []int{val, timestamp}
2. Check if size == cap, if yes, remove least used(oldest timestamp)
3. Add new Item in the map, then insert into minheap


Get(key)
1. check if key exist
2.




*/

type LRUCacheHeap struct {
}
