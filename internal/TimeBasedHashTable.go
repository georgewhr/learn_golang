package internal

/*

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".
Time based hash table

hashtable
1. pick up a right key, the key is the key, and maps to multiple entry, so it's 1:many mapping
2. solve collsion
3. same value same time? no overlapping time/value for same key

set(key, val, time), the insertion is always increasing with timestamp

ds1:
{
	"key1":{
		"t1":[val1, val2],
		"t2":[val1, val2],
		"t3":[val1, val2]
	}
}

ds2:
{
	"key1":[{t1, val1},{t1, val2}, {t2, val1}],
	"key2":[{t1, val1},{t1, val2}, {t2, val1}]
}


*/

// var INIT_SIZE = 100

type ShardKey struct {
	key  string
	time int
}

// type DataNode struct {
// 	val       string
// 	timestamp uint64
// }

type TBHT struct {
	Map map[ShardKey][]string
}

func ConstructionTBHT() *TBHT {
	return &TBHT{Map: make(map[ShardKey][]string)}
}

func (t *TBHT) Set(key string, val string, timeStamp int) {

	shardKey := ShardKey{key, timeStamp}
	if _, ok := t.Map[shardKey]; !ok {
		t.Map[shardKey] = []string{}
	}
	t.Map[shardKey] = append(t.Map[shardKey], val)

}

/*

it returns the value associated with the largest timestamp_prev
*/
// func (t *TBHT) Get(key string, time int) string {
// 	shardKey := ShardKey{key, time}
// 	if _, ok := t.Map[shardKey]; !ok {
// 		return ""
// 	} else {
// 		for _, val := range t.Map[shardKey]{
// 			if val ==
// 		}

// 	}
// }

// func serachVal(data []DataNode, time int) {
// 	low, high := 0, len(data)-1

// 	for low < high {
// 		mid := low + (high-low)/2
// 	}

// }
