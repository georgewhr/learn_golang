package internal

const MAP_SIZE = 50

type HashNode struct {
	key  int
	val  int
	Next *HashNode
}

type HashMap struct {
	BackingArr []*HashNode
}

func InitHashMap() *HashMap {
	return &HashMap{BackingArr: make([]*HashNode, MAP_SIZE)}
}

func (h *HashMap) InsertHashMap(key int, value int) {
	index := getIndex(key)

	if h.BackingArr[index] == nil {

	}

}

func getIndex(key int) (index int) {
	return int(hash(string(key))) % MAP_SIZE
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
