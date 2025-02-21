package internal

/*
Implement a LRU cache
Requirements:
Has Put, Get
discard least used items
each item will have an age that indicates the fequent of usage
both Put and Get will increase by 1


Get(k), get value of key, if not return -1
Put, Put k:v pair, return the existing value, otherwise return null




*/

type LRUCache struct {
}

// func Constructor(capacity int) LRUCache {
// 	return nil
// }

/*
need to be O(1)
if contains key, update exising key value pair

if not contains key, update

if map size greater than capacity, remove tail, then add new value in the head


*/

func (this *LRUCache) put(key int, value int) int {
	return 1
}

/*
need to be O(1)
if contains key, get value
if size == cap:
remove tail, add value

if not contains key, return -1




*/

func (this *LRUCache) get(key int) int {

	return 1
}
