package internal

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

/*
Design a Hashset
add(key)
contains(key)
remove(key)

[]bool


*/

type HashSet struct {
	backingArr []bool
	size       int
	cap        int
}

func InitHashSet(size int, cap int) *HashSet {
	return &HashSet{size: size, cap: cap}
}

func (h *HashSet) add(key int) bool {

	if h.contains(key) {
		return false
	} else {
		h.backingArr[hashFunc(key)] = true
	}

}

func (h *HashSet) remove(key int) {

}

func (h *HashSet) contains(key int) bool {
	if h.backingArr[hashFunc(key)] == true {
		return true
	} else {
		return false
	}

}

func (h *HashSet) hashFunc(v interface{}) int {
	hash := fnv.New32a()

	switch v := v.(type) {
	case int:
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v))
		hash.Write(b)
	case string:
		hash.Write([]byte(v))
	case float64:
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v))
		hash.Write(b)
	case bool:
		if v {
			hash.Write([]byte{1})
		} else {
			hash.Write([]byte{0})
		}
	default:
		// Convert other types to string as a fallback
		hash.Write([]byte(fmt.Sprintf("%v", v)))
	}

	return int(hash.Sum32()) % h.cap
}
