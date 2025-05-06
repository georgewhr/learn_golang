package internal

/*
Get(k), O(1)
Put(k, v),O(1)

1. Need a hashtable to store key -> value pair
2. Need a hashtable to store key -> frequency
3. Need a variable to store the least frequency
4. Need a hashtable to store frequeny to value(list of node to quickly serach for 1, 2)

Get()
check if key exist or not
IncreaseFreq(k)
return k_val[key]


Put()

if cap == size {
	RemoveMinFrequeny()
}

if key exist
	update valie
	IncreaseFreq(k)

else
	k_val[k] = val
	IncreaseFreq(k)
	this.minFreq = 1


RemoveMinFrequeny()
{
	node=freq_k[this.minFreq ]
	delete the first one(FIFO)

}

//This will be learn based on number of values in that freq
IncreaseFreq()

*/

type LFUCache struct {
	size    int
	kv      map[int]int
	kf      map[int]int
	fq      map[int][]int
	minFreq int
}

func LFUCacheInit() *LFUCache {
	return &LFUCache{size: 0, kv: map[int]int{}, kf: map[int]int{}, fq: map[int][]int{}}
}

func (c *LFUCache) Put(key int, val int) {

	if c.size == len(c.kv) {
		c.RemoveMinFrequeny()
	}

	if _, ok := c.kv[key]; ok {
		c.kv[key] = val
		c.IncreaseFreq(key)

	} else {
		c.kv[key] = val
		c.IncreaseFreq(key)
		c.minFreq = key
	}

}

func (c *LFUCache) Get(key int) int {
	if val, ok := c.kv[key]; !ok {
		return -1
	} else {
		c.IncreaseFreq(key)
		return val
	}
}

func (c *LFUCache) IncreaseFreq(key int) {
	olfFreq := c.kf[key]

	node := c.fq[olfFreq]

	for i, val := range node {
		if val == key {
			node = append(node[:i], node[i+1:]...)
			break
		}

	}

	if len(node) == 0 {
		delete(c.fq, olfFreq)
		if olfFreq == c.minFreq {
			c.minFreq++
		}

	}
	c.kf[key]++
	newFreq := c.kf[key]

	c.fq[newFreq] = append(c.fq[newFreq], key)

}
func (c *LFUCache) RemoveMinFrequeny() {
	minNodeKey := c.fq[c.minFreq][0]
	c.fq[c.minFreq] = c.fq[c.minFreq][1:]
	delete(c.kf, minNodeKey)
	delete(c.kv, minNodeKey)

}
