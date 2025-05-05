package internal

import "sort"

/*
题目是实现一个简易版的员工系统， 分4个level， 每个完成了才能解锁下一层。
lvl 1. 写两个功能, 一个是添加新员工信息包括姓名， 职位和薪资（时薪）。 另一个是上下班打卡。
lvl 2. 支持查询某个员工的在办公室的总时间。
以下内容需要积分高于 188 您已经可以浏览
按时间找出top K员工
lvl 3. 支持升职加薪功能， 先announce升职， 然后在之后的第一次进入office时生效。支持查询任意时间段的总薪酬（打卡时间 * 时薪）。
lvl 4. 圈定某些时间段为双倍薪资。

*/

type Employee struct {
	workerId         string
	position         string
	compensation     string
	totalTime        int
	currentStatus    int // 0 init, 1 in office, 2 left office
	currentBadgeIn   string
	currentBadgeOut  string
	schedule         []*(WorkTime)
	name             string
	promoteStatus    string
	promoteStartTime string
	promotePosition  string
	promoteComp      string
}

type WorkTime struct {
	badgeIn  string
	badgeOut string
	pay      string
	position string
}

type CompanyWorker struct {
	worker map[string]*Employee
}

func InitCompanyWorker() *CompanyWorker {
	return &CompanyWorker{worker: make(map[string]*Employee)}
}

func (this *CompanyWorker) AddWorker(workerId string, position string, compensation string) bool {
	this.worker[workerId] = &Employee{workerId: workerId, position: position, compensation: compensation}
	return true
}

func (this *CompanyWorker) Register(workerId string, time string) bool {

	if this.worker[workerId].currentStatus != 1 {
		this.worker[workerId].currentBadgeIn = time
		this.worker[workerId].currentStatus = 1

		// if this.worker[workerId].promotePosition == "PENDING" && convertStrInt(time) >= convertStrInt(this.worker[workerId].promoteStartTime) {
		// 	this.worker[workerId].promotePosition = "DONE"
		// 	this.worker[workerId].compensation = this.worker[workerId].promoteComp
		// 	this.worker[workerId].totalTime = 0
		// }

	} else if this.worker[workerId].currentStatus == 1 {
		this.worker[workerId].currentBadgeOut = time
		this.worker[workerId].currentStatus = 2

		newWorkTime := &WorkTime{badgeIn: this.worker[workerId].currentBadgeIn, badgeOut: this.worker[workerId].currentBadgeOut, pay: this.worker[workerId].compensation, position: this.worker[workerId].position}
		this.worker[workerId].schedule = append(this.worker[workerId].schedule, newWorkTime)
		this.worker[workerId].totalTime += convertStrInt(this.worker[workerId].currentBadgeOut) - convertStrInt(this.worker[workerId].currentBadgeIn)

		if this.worker[workerId].promoteStatus == "PENDING" && convertStrInt(time) >= convertStrInt(this.worker[workerId].promoteStartTime) {
			this.worker[workerId].position = this.worker[workerId].promotePosition
			this.worker[workerId].compensation = this.worker[workerId].promoteComp
			this.worker[workerId].totalTime = 0

		}

	}

	return true

}

// func (this *CompanyWorker) Get(workerId string) bool {

// 	for _, val := range this.worker[workerId].schedule {
// 		(val.badgeOut-val.badgeIn) * val.pay
// 	}

// }

func (this *CompanyWorker) TopN(n string, position string) string {

	tempList := make([]*Employee, 0)

	ret := ""

	for _, val := range this.worker {
		if val.position == position {
			tempList = append(tempList, val)
		}

	}

	sort.Slice(tempList, func(i, j int) bool {
		if tempList[i].totalTime == tempList[j].totalTime {
			return tempList[i].name < tempList[j].name
		}
		return tempList[i].totalTime > tempList[j].totalTime
	})

	// var parts []string

	if convertStrInt(n) >= len(tempList) {
		for _, val := range tempList {

			ret += (val.workerId + convertIntStr(val.totalTime) + ",")

		}
	} else {
		for _, val := range tempList[:convertStrInt(n)] {

			ret += val.workerId + convertIntStr(val.totalTime)

		}
	}

	return ""

}

func (this *CompanyWorker) Promote(workerId string, newPosition string, newComp string, startTime string) bool {
	if this.worker[workerId].promoteStatus == "PENDING" {
		return false
	}

	this.worker[workerId].promoteStatus = "PENDING"
	this.worker[workerId].promotePosition = newPosition
	this.worker[workerId].promoteComp = newComp
	this.worker[workerId].promoteStartTime = startTime

	return true

}

func (this *CompanyWorker) CalculateSalary(workerId string, startTime int, endTime int) bool {

	totalSal := 0
	for _, val := range this.worker[workerId].schedule {

		if convertStrInt(val.badgeIn) >= startTime && convertStrInt(val.badgeOut) <= endTime {
			totalSal += (convertStrInt(val.badgeOut) - convertStrInt(val.badgeIn)) * convertStrInt(val.pay)
		}

	}

	return true

}
