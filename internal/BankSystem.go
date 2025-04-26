package internal

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





*/
