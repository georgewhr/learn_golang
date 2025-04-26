package internal

import (
	"fmt"
	"sort"
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
	db       map[KeySet][]*Val
	backUpDB map[string]map[KeySet][]*Val
}

func InitInDB() *InMemoryDB {
	return &InMemoryDB{db: make(map[KeySet][]*Val), backUpDB: make(map[string]map[KeySet][]*Val)}
}

func (this *InMemoryDB) Set(key string, field string, val string) string {
	k := KeySet{key: key, field: field}

	if _, ok := this.db[k]; !ok {
		this.db[k] = append(this.db[k], &Val{value: val})
	} else {
		this.db[k][0].value = val
	}

	return ""

}

func (this *InMemoryDB) Get(key string, field string) string {

	k := KeySet{key: key, field: field}
	if _, ok := this.db[k]; ok {
		return this.db[k][0].value
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
			for _, item := range val {
				res = res + fmt.Sprintf("%s(%s), ", k.field, item.value)
			}

		}
	}

	return res

}
func (this *InMemoryDB) ScanPrefix(key string, field string) string {

	res := ""
	for k, val := range this.db {
		if key == k.key && field == k.field {
			res = res + fmt.Sprintf("%s(%s), ", k.key, val[0].value)
		}
	}

	return res

}

func (this *InMemoryDB) SetAt(key string, field string, val string, time string) string {

	// this.Set(key, field, val)
	k := KeySet{key: key, field: field}
	this.db[k] = append(this.db[k], &Val{value: val, timeStamp: convertStrInt(time)})
	return ""
}

func (this *InMemoryDB) SetAtWithTTL(key string, field string, val string, time string, ttl string) string {

	k := KeySet{key: key, field: field}
	this.db[k] = append(this.db[k], &Val{value: val,
		timeStamp: convertStrInt(time),
		ttl:       convertStrInt(ttl)})
	return ""
}

func (this *InMemoryDB) DeleteAt(key string, field string, time string) bool {

	k := KeySet{key: key, field: field}

	if _, ok := this.db[k]; !ok {
		return false
	}

	for i, val := range this.db[k] {
		if val.ttl > convertStrInt(time) && val.timeStamp == convertStrInt(time) {
			this.db[k] = append(this.db[k][:i], this.db[k][:i+1]...)
			return true
		}

	}

	return false

}

func (this *InMemoryDB) GetAt(key string, field string, time string) string {
	k := KeySet{key: key, field: field}

	res := ""

	if _, ok := this.db[k]; !ok {
		return ""
	}
	for _, val := range this.db[k] {
		if val.ttl > convertStrInt(time) && val.timeStamp == convertStrInt(time) {
			res = res + fmt.Sprintf("%s(%s), ", k.field, val.value)
		}

	}

	return res

}

// func (this *InMemoryDB) ScanAt(key string, time string) string {

// 	res := ""
// 	for k, v := range this.db {

// 		if k.key == key && v.ttl > convertStrInt(time) {
// 			res = res + fmt.Sprintf("%s(%s), ", k.field, v.value)
// 		}

// 	}

// 	return res

// }

func (this *InMemoryDB) Backup(time string) string {

	// this.backUpDB = make(map[string]map[KeySet]*Val)
	res := 0
	this.backUpDB[time] = make(map[KeySet][]*Val)
	for k, v := range this.db {

		for _, item := range v {
			if item.ttl > convertStrInt(time) {

				newItem := &Val{value: item.value, ttl: item.ttl, timeStamp: item.timeStamp}
				newItem.ttl = newItem.ttl - convertStrInt(time) + 1
				this.backUpDB[time][k] = make([]*Val, 0)
				this.backUpDB[time][k] = append(this.backUpDB[time][k], newItem)
				res++
			}
		}
	}

	return convertIntStr(res)

}

func (this *InMemoryDB) Restore(time string, timeToRestore string) string {

	// this.backUpDB = make(map[KeySet]*Val)
	res := ""
	timeList := []int{}
	for k, _ := range this.backUpDB {
		timeList = append(timeList, convertStrInt(k))
	}

	sort.Slice(timeList, func(i, j int) bool {
		return timeList[i] > timeList[j]
	})

	var latestBackupTime int
	for _, v := range timeList {
		if v <= convertStrInt(timeToRestore) {
			latestBackupTime = v
			break
		}
	}

	tempDb := this.backUpDB[convertIntStr(latestBackupTime)]
	this.db = tempDb
	for k, _ := range this.db {

		for index, _ := range this.db[k] {
			this.db[k][index].ttl = this.db[k][index].ttl + convertStrInt(time)

		}
	}

	return res

}

//
