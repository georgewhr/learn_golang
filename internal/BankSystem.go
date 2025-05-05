package internal

import (
	"fmt"
	"sort"
)

/*


刚做完Anthropic OA，拿到接近满分，差一个没完成
考题是banking system，没来得及截图，但稍微还有些记忆分享给大家
1. create_account, deposit, pay: 基本帐户操作，只需要使用dict储存account name & amount，稍微检查一下重複即可，大约花5分钟做完
create_account(self, timestamp: int, account_id: str) -> bool
deposit(self, timestamp: int, account_id: str, amount: int) -> int | None
pay(self, timestamp: int, account_id: str, amount: int) -> int | None
2. 求 top N activities: 排序总operation加起来最高金额的用户，需要先用amount desc sort，接着再用account name asc sort，这要注意一下，我大约10分钟搞定
top_activity(self, timestamp: int, n: int) -> list[str]
3. transfer, accept_transfer:


1. record and hold transactions - deposits and transfers
2. do data metrics returning topK accounts with outgoing money
3. Add scheduled transactions and cancel them
4. merge two accounts while maintaining separate account histories)


Level 1
Initially, the banking system does not contain any accounts, so implement operations to allow account creation, deposits, and transfers between different accounts:create_account(self, timestamp: int, account_id: str) -> bool
— should create a new account with the given identifier if it doesn’t already exist. Returns True if the account was successfully created or False if an account with that id already exists.deposit(self, timestamp: int, account_id: str, amount: int) -> int | None
— should deposit the given amount of money to the specified account id. Returns the balance of the account after the operation has been processed or None if the specified account doesn’t exist. The operation has been processed before the result is returned.transfer(self, timestamp: int, source_account_id: str, target_account_id: str, amount: int) -> int | None
— should transfer the given amount of money from account source_account_id to account target_account_id. Returns the balance of source_account_id if the transfer was successful or None otherwise.
Returns None if source_account_id or target_account_id doesn’t exist.
Returns None if source_account_id and target_account_id are the same.
Returns None if account source_account_id has insufficient funds to perform the transfer.

Level 2
The bank wants to identify people who are not keeping money in their accounts, so implement operations to support ranking accounts based on outgoing transactions:top_spenders(self, timestamp: int, n: int) -> list[str]
— should return the identifiers of the top n accounts with the highest outgoing transactions — the total amount of money either transferred out of or paid/withdrawn (the pay operation will be introduced in level 3) — sorted in descending order, or in case of a tie, sorted alphabetically by account_id in ascending order. The result should be a list of strings in the following format:
["account_id_1>(total_outgoing_1)", "account_id_2>(total_outgoing_2)", ..., "account_id_n>(total_outgoing_n)"]
If less than n accounts exist in the system, then return all their identifiers (in the described format).
Cashback (an operation that will be introduced in level 3) should not be reflected in the calculations for total outgoing transactions.

Level 3
The banking system should allow scheduling payments with some cashback and checking the status of scheduled payments:pay(self, timestamp: int, account_id: str, amount: int) -> str | None
— should withdraw the given amount of money from the specified account.
All withdraw transactions provide a 2% cashback. The cashback amount (rounded down to the nearest integer) will be refunded to the account 24 hours after the withdrawal, if the withdrawal was successful. If the refund is successful (i.e., the account hasn’t been removed from the system), return the given amount; otherwise, return None. Returns a unique identifier for the payment transaction (e.g., "payment1", "payment2", etc.).
Additional conditions:
Returns None if account_id doesn’t exist.
Returns None if account_id has insufficient funds to perform the payment.
top_spenders should now also account for the total amount of money withdrawn from accounts.
The waiting period for cashback is 24 hours, equal to 24 * 60 * 60 * 1000 = 86400000 milliseconds (the unit for timestamps).
So, cashback will be processed at timestamp + 86400000.get_payment_status(self, timestamp: int, account_id: str, payment: str) -> str | None
— should return the status of the payment transaction for the given payment. Specifically:
Returns None if account_id doesn’t exist.
Returns None if the given payment doesn’t exist for the specified account.
Returns None if the payment transaction was for an account with a different identifier from account_id.
Returns a string representing the payment status: "IN_PROGRESS" or "CASHBACK_RECEIVED".

Level 4
The banking system should support merging two accounts while retaining both accounts’ balance and transaction histories.merge_accounts(self, timestamp: int, account_id_1: str, account_id_2: str) -> bool
— should merge account_id_2 into account_id_1. Returns True if accounts were successfully merged, or False otherwise. Specifically:
Returns False if account_id_1 is equal to account_id_2.
Returns False if account_id_1 or account_id_2 doesn’t exist.
All pending cashback refunds for account_id_2 should still be processed, but refunded to account_id_1 instead.
After the merge, it must be possible to check the status of payment transactions in account_id_2 with payment identifiers by replacing account_id_2 with account_id_1.
The balance of account_id_2 should be added to the balance for account_id_1.
top_spenders operations should recognize merged accounts — the total outgoing transactions for merged accounts should be the sum of all money transferred and/or withdrawn in both accounts.
account_id_2 should be removed from the system after the merge.
get_balance(self, timestamp: int, account_id: str, time_at: int) -> int | None
— should return the total amount of money in the account account_id at the given timestamp time_at. If the specified account did not exist at a given time time_at, returns None.
If queries have been processed at timestamp time_at, get_balance must reflect the account balance after the query has been processed.
If the account was merged into another account, the merged account should inherit its balance history

*/

type AcountInfo struct {
	accountId          string
	balance            int
	createdTime        int
	transactions       int
	transactionsAmount int
	outGoingAmount     int
	moneyWithDrawed    int
	payMentInfo        map[string]*Payment
}

type Payment struct {
	paymentId string
	amount    int
	status    string
	time      string
	position  string
}

type Transfer struct {
	srcAccount    string
	targetAccount string
	time          string
	amount        int
	status        string
}

type BankingSystem struct {
	bkSystem map[string]*AcountInfo
}

func InitBankingSystem() *BankingSystem {
	return &BankingSystem{bkSystem: make(map[string]*AcountInfo)}
}

func (this *BankingSystem) CreateAccount(timeStamp int, accountId string) bool {
	if _, ok := this.bkSystem[accountId]; !ok {
		this.bkSystem[accountId] = &AcountInfo{accountId: accountId, createdTime: timeStamp, balance: 0, transactions: 0, transactionsAmount: 0}
		return true
	} else {
		return false
	}
}

func (this *BankingSystem) Deposit(timeStamp int, accountId string, amount int) bool {
	if _, ok := this.bkSystem[accountId]; !ok {
		return false
	} else {
		this.bkSystem[accountId].transactionsAmount += amount
		this.bkSystem[accountId].balance = this.bkSystem[accountId].balance + amount
		return true
	}
}

func (this *BankingSystem) Pay(timeStamp string, accountId string, amount int) bool {

	if _, ok := this.bkSystem[accountId]; !ok {
		return false
	}

	if amount > this.bkSystem[accountId].balance {
		return false
	}

	if this.bkSystem[accountId].payMentInfo == nil {
		this.bkSystem[accountId].payMentInfo = make(map[string]*Payment)
		fmt.Print("sasda")
	}
	paymengId := accountId + "_" + convertIntStr(this.bkSystem[accountId].transactions)
	this.bkSystem[accountId].payMentInfo[paymengId] = &Payment{paymentId: paymengId, time: timeStamp, amount: amount, status: "IN_PROGRESS"}
	this.bkSystem[accountId].transactions++
	this.bkSystem[accountId].balance -= amount
	this.bkSystem[accountId].transactionsAmount += amount

	return true

}

func (this *BankingSystem) GetPaymentStatus(timeStamp int, accountId string, paymentId string) bool {

	if convertStrInt(this.bkSystem[accountId].payMentInfo[paymentId].time)+86400000 <= timeStamp {
		this.bkSystem[accountId].payMentInfo[paymentId].status = "CASHBACK_RECEIVED"
		this.bkSystem[accountId].balance = this.bkSystem[accountId].balance*2/100 + this.bkSystem[accountId].balance
		return true
	}

	return false

}
func (this *BankingSystem) Transfer(timeStamp int, srcAccountId string, targetAccountId string, amount int) bool {
	if _, ok := this.bkSystem[srcAccountId]; !ok {
		return false
	}

	if _, ok := this.bkSystem[targetAccountId]; !ok {
		return false
	}

	if srcAccountId == targetAccountId {
		return false
	}

	if amount < 0 {
		return false
	}

	tempVal := this.bkSystem[srcAccountId].balance - amount

	if tempVal < 0 {
		return false
	}
	this.bkSystem[srcAccountId].balance = tempVal
	this.bkSystem[targetAccountId].balance += tempVal
	return true

}

func (this *BankingSystem) Withdraw(timeStamp int, accountId string, amount int) bool {
	if _, ok := this.bkSystem[accountId]; !ok {
		return false
	} else {
		tempVal := this.bkSystem[accountId].balance - amount
		if tempVal < 0 {
			return false
		}
		this.bkSystem[accountId].balance = tempVal
		return true
	}
}

func (this *BankingSystem) TopN(n int) string {

	tempList := make([]*AcountInfo, 0)
	res := ""

	for _, val := range this.bkSystem {
		tempList = append(tempList, val)
	}

	// if tie, sort by account ID
	sort.Slice(tempList, func(i, j int) bool {
		if tempList[i].transactions == tempList[j].transactions {
			return tempList[i].accountId < tempList[j].accountId
		}
		return tempList[i].transactions > tempList[j].transactions
	})

	sort.Slice(tempList, func(i, j int) bool {
		return tempList[i].transactions > tempList[j].transactions
	})

	tempList = tempList[:n]

	// sort alphabetically
	sort.Slice(tempList, func(i, j int) bool {
		return tempList[i].accountId < tempList[j].accountId
	})

	for _, val := range tempList {
		res = res + val.accountId + " "
	}
	return res

}

func (this *BankingSystem) MergeAccount(account1 string, account2 string) string {

	this.bkSystem[account1].balance = this.bkSystem[account1].balance + this.bkSystem[account2].balance
	this.bkSystem[account2].balance = 0
	for k, v := range this.bkSystem[account2].payMentInfo {
		this.bkSystem[account1].payMentInfo[k] = &Payment{paymentId: v.paymentId, amount: v.amount, status: v.status, time: v.time}
	}

	delete(this.bkSystem, account2)

	return ""

}
