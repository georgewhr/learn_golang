package internal

import (
	"fmt"
	"sort"
	"strings"
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
	db       map[KeySet]*Val
	backUpDB map[string]map[KeySet]*Val
}

func InitInDB() *InMemoryDB {
	return &InMemoryDB{db: make(map[KeySet]*Val), backUpDB: make(map[string]map[KeySet]*Val)}
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
	tempRes := make([][]string, 0)
	for k, val := range this.db {
		// ksplit := strings.Split(k, "_")
		if k.key == key {
			tempRes = append(tempRes, []string{k.field, val.value})
			// for _, item := range val {
			// res = res + fmt.Sprintf("%s(%s), ", k.field, val.value)
			// }

		}
	}

	if len(tempRes) == 0 {
		return ""
	}

	sort.Slice(tempRes, func(i, j int) bool {
		return tempRes[i][0] > tempRes[j][0]
	})

	var parts []string

	for _, pair := range tempRes {
		parts = append(parts, fmt.Sprintf("%s(%s)", pair[0], pair[1]))

	}

	result := strings.Join(parts, ", ")
	fmt.Println(result)

	// for _, val := range tempRes {
	// 	res = res + fmt.Sprintf("%s(%s), ", val[0], val[1])

	// }

	return res

}
func (this *InMemoryDB) ScanPrefix(key string, prefix string) string {

	// res := ""
	tempRes := make([][]string, 0)
	for k, val := range this.db {
		if k.key == key && strings.HasPrefix(val.value, prefix) {
			tempRes = append(tempRes, []string{k.field, val.value})

		}

	}

	if len(tempRes) == 0 {
		return ""
	}

	sort.Slice(tempRes, func(i, j int) bool {
		return tempRes[i][0] > tempRes[j][0]
	})

	var parts []string

	for _, pair := range tempRes {
		parts = append(parts, fmt.Sprintf("%s(%s)", pair[0], pair[1]))

	}

	result := strings.Join(parts, ", ")
	fmt.Println(result)

	return result

}

func (this *InMemoryDB) SetAt(key string, field string, val string, time string) string {

	// this.Set(key, field, val)
	k := KeySet{key: key, field: field}

	if _, ok := this.db[k]; !ok {
		this.Set(key, field, val)
	}

	this.db[k].timeStamp = convertStrInt(time)

	return ""
}

func (this *InMemoryDB) SetAtWithTTL(key string, field string, val string, time string, ttl string) string {

	k := KeySet{key: key, field: field}
	if _, ok := this.db[k]; !ok {
		this.Set(key, field, val)
	}

	this.db[k].timeStamp = convertStrInt(time)
	this.db[k].ttl = convertStrInt(ttl) + convertStrInt(time)
	return ""
}

func (this *InMemoryDB) DeleteAt(key string, field string, time string) bool {

	k := KeySet{key: key, field: field}

	if _, ok := this.db[k]; !ok {
		return false
	}

	if this.db[k].ttl < convertStrInt(time) {
		delete(this.db, k)
	} else {
		return false
	}

	return false

}

func (this *InMemoryDB) GetAt(key string, field string, time string) string {
	k := KeySet{key: key, field: field}

	res := ""

	if _, ok := this.db[k]; !ok {
		return ""
	}

	if this.db[k].ttl > convertStrInt(time) {
		return this.db[k].value
	}

	return res

}

func (this *InMemoryDB) ScanAt(key string, time string) string {

	// res := ""
	tempRes := make([][]string, 0)
	for k, v := range this.db {

		if k.key == key && v.ttl > convertStrInt(time) {
			tempRes = append(tempRes, []string{k.field, v.value})
		}

	}

	if len(tempRes) == 0 {
		return ""
	}

	sort.Slice(tempRes, func(i, j int) bool {
		return tempRes[i][0] > tempRes[j][0]
	})

	var parts []string

	for _, pair := range tempRes {
		parts = append(parts, fmt.Sprintf("%s(%s)", pair[0], pair[1]))

	}

	result := strings.Join(parts, ", ")
	fmt.Println(result)

	return result

}

func (this *InMemoryDB) Backup(time string) string {

	res := 0
	this.backUpDB[time] = make(map[KeySet]*Val)
	for k, v := range this.db {

		// for _, item := range v {
		if v.ttl > convertStrInt(time) {

			newItem := &Val{value: v.value, ttl: v.ttl, timeStamp: v.timeStamp}
			newItem.ttl = newItem.ttl - convertStrInt(time) + 1
			this.backUpDB[time][k] = newItem
			// this.backUpDB[time][k] = append(this.backUpDB[time][k], newItem)
			res++
		}
		// }
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

	this.db = map[KeySet]*Val{}

	tempDb := this.backUpDB[convertIntStr(latestBackupTime)]
	// this.db = tempDb
	for k, _ := range tempDb {
		this.db[k] = &Val{}
	}

	return res

}

//
