package internal

import (
	"fmt"
)

/*
写一个in memory db
in memory db implementation, store key record pairs, the record would be in a format of field, value
level 1  - set, get, compare_and_set, compare_and_delete
level 2 - scan the record for the given key, output a string formatted results
level 3 - support timestamp and ttl of the field value, each field value could be expire after ttl
level 4 - get a historical value
时间比较紧 debug不是很方便 提前准备下会比较好 我准备了bank的题结果遇到了这个

*/

type Val struct {
	value     string
	timeStamp int
	ttl       int
}

type KeySet struct {
	key   string
	field string
}

type InMemoryDB struct {
	db map[KeySet]*Val
}

func InitInDB() *InMemoryDB {
	return &InMemoryDB{db: make(map[KeySet]*Val)}
}

func (this *InMemoryDB) Set(key string, field string, val string) string {
	k := KeySet{key: key, field: field}

	if _, ok := this.db[k]; !ok {
		this.db[k] = &Val{value: val}
	} else {
		this.db[k].value = val
	}

	return ""

}

func (this *InMemoryDB) Get(key string, field string) string {

	k := KeySet{key: key, field: field}
	if _, ok := this.db[k]; ok {
		return this.db[k].value
	} else {
		return ""
	}

}

func (this *InMemoryDB) Delete(key string, field string) bool {

	k := KeySet{key: key, field: field}
	if _, ok := this.db[k]; ok {
		delete(this.db, k)
		return true
	} else {
		return false
	}
}

func (this *InMemoryDB) Scan(key string) string {

	res := ""
	for k, val := range this.db {
		// ksplit := strings.Split(k, "_")
		if k.key == key {
			res = res + fmt.Sprintf("%s(%s), ", key, val.value)
		}
	}

	return res

}
func (this *InMemoryDB) ScanPrefix(key string, field string) string {

	res := ""
	for k, val := range this.db {
		if key == k.key && field == k.field {
			res = res + fmt.Sprintf("%s(%s), ", k.key, val.value)
		}
	}

	return res

}

func (this *InMemoryDB) SetAt(key string, field string, val string, time string) string {

	this.Set(key, field, val)
	k := KeySet{key: key, field: field}
	this.db[k].timeStamp = convertStrInt(time)
	return ""
}

func (this *InMemoryDB) SetAtWithTTL(key string, field string, val string, time string, ttl string) string {

	this.Set(key, field, val)
	k := KeySet{key: key, field: field}
	this.db[k].timeStamp = convertStrInt(time)
	this.db[k].ttl = convertStrInt(time) + convertStrInt(ttl)
	return ""
}
