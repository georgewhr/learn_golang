package main

import (
	"container/heap"
	"fmt"
	"george/internal"
	"math"
	"sort"
)

type ListNode = internal.ListNode
type List = internal.LinkedList
type PriorityQueue = internal.PriorityQueue
type Stack = internal.Stack
type TreeNode = internal.TreeNode
type TBHS = internal.TBHT
type Test struct {
	name string
	size int
}

func main() {

	ttt := []Test{}
	ttt = append(ttt, Test{name: "As", size: 20})
	ttt = append(ttt, Test{name: "John", size: 50})
	ttt = append(ttt, Test{name: "Bs", size: 20})
	ttt = append(ttt, Test{name: "Jason", size: 50})

	sort.Slice(ttt, func(i, j int) bool {
		if ttt[i].size == ttt[j].size {
			return ttt[i].name < ttt[j].name
		}
		return ttt[i].size > ttt[j].size
	})

	fmt.Println("Hello, world.")

	employee := internal.InitCompanyWorker()
	employee.AddWorker("george", "J", "100")
	employee.AddWorker("vera", "J", "150")
	employee.Register("george", "10")
	employee.Register("george", "12")
	employee.Register("george", "11")
	employee.CalculateSalary("george", 10, 14)
	employee.TopN("10", "J")

	cloudStorage := internal.Init()
	inMemDb := internal.InitInDB()

	inMemDb.SetAtWithTTL("A", "T", "G", "3", "5")
	inMemDb.SetAtWithTTL("A", "B", "C", "4", "14")
	inMemDb.SetAtWithTTL("A", "D", "E", "5", "15")
	inMemDb.Backup("4")
	inMemDb.Backup("12")
	inMemDb.Restore("100", "12")

	// cloudStorage.Add("/root/file1.mp3", "10")
	// cloudStorage.Add("/root/file2.mp3", "56")
	// cloudStorage.Add("/dir100/file100.mp3", "1")
	// cloudStorage.Copy("/dir100/file100.mp3", "/dir100/copy/file100.mp3")
	// cloudStorage.Add("/dir/file2.txt", "20")
	// cloudStorage.Add("/dir3/file3.txt", "50")

	// cloudStorage.FindFile("/root", ".mp3")
	cloudStorage.AddUsers("george", "100")
	cloudStorage.AddFileBy("george", "/dir/file2.txt", "40")
	cloudStorage.Copy("/dir/file2.txt", "/dircopy/file2.tx")
	cloudStorage.GetFileSzie("/dir/file2.txt")

	// cloudStorage.AddUsers("vera", "200")
	cloudStorage.AddFileBy("george", "/dir101/file100.mp3", "40")
	cloudStorage.AddFileBy("george", "/dir101/file1001.mp3", "50")
	cloudStorage.Copy("/dir101/file1001.mp3", "/dir100/copy1/file100.mp3")
	cloudStorage.UpdateCapacity("george", "50")

	bank := internal.InitBankingSystem()
	bank.CreateAccount(100, "george")
	bank.CreateAccount(120, "vera")
	bank.Deposit(130, "george", 100)
	bank.Deposit(130, "vera", 200)
	bank.Pay("200", "george", 20)
	bank.Pay("201", "vera", 15)
	bank.Pay("202", "george", 40)
	bank.Transfer(300, "george", "vera", 5)
	bank.Withdraw(400, "george", 10)
	bank.GetPaymentStatus(86500000, "george", "george_0")

	// cloudStorage.AddFileBy("george", "/dir102/file1001.mp3", "50")

	// test := []string{"cir", "car"}
	// test := []string{"fla", "flw"}
	// test1 := "abc"
	// fmt.Printf("number is %s", test1[0:1])
	// output := longestCommonPrefix(test)
	// fmt.Printf("number is %f\n", output)
	basicImplementations()
	myList := List{}
	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)
	myList.Insert(4)
	myList.Insert(5)

	// myNode := removeNthNodeFromEndofList(myList.Head, 2)
	// fmt.Printf("number is %f", myNode.Val)

	var georgetest = []int{2, 0, 2, 1, 1, 0}
	sortColors(georgetest)

	// b := []int{2, 7, 11, 13}
	// nums := twoSum(b, 9)
	// fmt.Println(nums)

	// sortedA := []int{1, 2}
	// sortedB := []int{3, 4}

	// newArr := findMedianSortedArrays(sortedA, sortedB)
	// fmt.Printf("number is %f", newArr)

	// testStr := longestPalindromeFromMiddle("babc")
	// fmt.Printf("number is %s", testStr)

	// testNum := reverse(1534236469)
	// fmt.Printf("number is %d", testNum)

	// isPalindrome(123)

	// a := &b
	// fmt.Println("asdasd %d", (*a)[0])

	// newListNode := &ListNode{}
	// for aa := 0; aa < 3; aa++ {
	// 	newListNode = &ListNode{Val: aa}
	// 	newListNode = newListNode.Next
	// }

	myList1 := List{}
	myList1.Insert(4)
	myList1.Insert(3)
	myList1.Insert(4)
	myList1.Insert(1)
	// myList1.Insert(4)
	// myList1.Insert(5)

	result := print_list_recursive(myList1.Head, -1)
	fmt.Printf("result is %v", result)

	myTBHS := internal.ConstructionTBHT()
	myTBHS.Set("key1", "val1", 123)

	// myList2 := List{}
	// myList2.Insert(6)
	// myList2.Insert(4)
	// myList2.Insert(2)
	// newHead := mergedSortedList(myList1.Head, myList2.Head)
	// fmt.Println(newHead)

	newHead := rotateRight(myList1.Head, 2)
	fmt.Println(newHead)
	// myList.PrintList()
	// myList1.PrintList()

	// myList2 := List{}
	// myList2.Head = add2List(myList.Head, myList1.Head)
	// myList2.PrintList()

	// fmt.Println("length is ", lengthOfLongestSubstring("pwwkew"))

	// myHashMap := internal.InitHashMap()

	// fmt.Println(myHashMap.BackingArr[0])

	myStack := internal.InitStack(2)
	myStack.Push('a')
	myStack.Push('b')

	myStack1 := internal.InitStack(3)
	myStack1.Push('a')
	myStack1.Push('b')

	myQ := internal.InitQueue(5)
	myQ.Add(1)
	myQ.Add(2)
	myQ.Add(3)
	myQ.Remove()
	myQ.Remove()
	// // myStack.Push(3)
	// myStack.PrintStack()
	// ss := myStack.Pop()
	// fmt.Printf("Successfully cast to rune: %c\n", ss.(rune))
	// myStack.PrintStack()
	// result := isValidParenthese("()")
	// fmt.Println(result)

	data := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}

	// Construct the list of ListNode pointers
	var lists []*ListNode
	for _, vals := range data {
		lists = append(lists, createList(vals))
	}

	for _, head := range lists {
		printList(head)
	}

	newList := mergeKLists2(lists)
	fmt.Println(*newList)

	sampleArr := [3]int{1, 1, 2}

	results := removeDuplicatedArr(sampleArr[:])
	fmt.Println(results)

	test100 := []int{1, 2, 3, 4}
	productExceptSelf(test100)

	lengthOfLongestSubstringWithoutRepeatingChar("dvdf")
	longestRepeatingChar("AABABBA", 1)

	ss := checkInclusion("trfhcbogglim", "amwfpqwfwkarvhfcisywzaahtbdbuicxmjseeoudwfcdxetbmacayfikolbdxkocezhalfhxabwvuddcyazwiqiwefgolzvrpdxcuskpsmwhslpeygjrvvajajafppcwkqhxwkigemfullbqkvuqwfnqnhxiltyfcpfdyumfwyelmtzbdccmbvxedgfimmsqwxmopvxmuonuzyzlhpeunailpydcqybghdwvqxrpautsvrhfxprdqlgzownvincoxjnjwrqrdgpegtgvlifbbautkfqbhqiftbmxadvorqjnqlsceuctazxgofmchypspqvwyzoeejqfkvvwftvagajafmafvytslubpzalnahjknarjswkxmzzlmlokrifiopjcamvynmmuegnzvveetssuucqclbzxgjwbsflyelpdsvzicdnlenuxggcsrckfdixsqcjrzsbztgvxbpktlbdqrcqoatgxqhwehqiuqjnldursyzplwlcdvwrmlknviqtexwgqovwbcdugdblakufxcapvkvhraacetowtcypfxlvwmwdafesfgqezspbvqzxicblrdsmmdzunpcmzvysgbnspuldkppwlrsrnnewwjquhzrodmsbgbycvrzmtnskyuqqoqkxyakojewbbtntbdjuncpgbwgrtvewyscyujnqtpuaulrnjqmdujxydwzfyqfnxmjqogibxqeuqdxsdjjpootpzmhcvoeyvdspktyjzadkfwcdulsuktottgpvptjfydvpdxoznzhbkmitaiywqklwrktmzsyndnqmtapaaadzkynfxiwqxtekcbkmcwhwwdylvoxosxcexeceavpfptdlkyinhdobrnxfdbtuomjojmzeytlntkundrydqmkiayounnbhfxrlriuatzumgfcyniicwhtsaffhnxamwjtgbxvewtgovvrjvblrlvrghyoiicgvyorzjgecmxqeiwpuubfrnkmpynwywqczqdpeinebgfyrhouvifthoaariadurpbrexbfnuwgkbmgowjuaysnmidptzetckscxvxttdogpywxdvaktmkispgyghfazxyxslyjhqorndzpjepmwiuisfhvacnpkthbfrasndrfkfuhpetlnfugmqhqpvtvlwumhxduxxmugstcbksvqholothhftzungtxdysudnixkzekpdlgddnvyfuitcedxvjfsjxhbcrenufafxzdrumeavumdbvvgpodgtsjzznxkpbfltchmogigordwcpcanomjznfmsxpzqgxigjpybooxsgyiuxskowkdpypnzpgebowqefomcpmfilixgzvoffvmcypgyrwhwaelfpclzaoldlaimtnszckziyqewrtewpfyhphxruytifwtodznvxmxwoibqvtmynpqshnmiymrayaenoiknjqzwoltqhaganjdwzkncathqrgcigaguimqgznupmsikurxjltfydqiozmddxydgtsvwoloqtlqhryfqmcsfetvtjkauyjgillobotqfhzsyjtcjsiqxhwoaucluagbltdwroayydlwzytpqlsxkbrgcavvaqvlggewskeflsejklqexjvcudzaanxrgnkwygokcuxkvypsh")
	fmt.Println(ss)

	var rootNode = &TreeNode{Val: 2}
	var leftNode = &TreeNode{Val: 1}
	var rightNode = &TreeNode{Val: 3}
	rootNode.Left = leftNode
	rootNode.Right = rightNode
	// isValidBST2(rootNode)
	// kSmallestBST(rootNode, 1)
	// generateParenthesis(3)
	// subsets([]int{1, 2, 3})

	// combinationSum2UniqueNumber([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	// combinationSum2UniqueNumber([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	// combinationSum2UniqueNumber([]int{1, 2, 2, 2, 5}, 5)
	// threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10})
	// coinChangeRec([]int{1, 5, 10}, 12)
	// combinationSumAny([]int{3, 3}, 6)
	// combinationSum4([]int{1, 2, 3}, 4)
	// comboSumShortrest([]int{2, 3, 5}, 8)
	// climbingStairsBC(3)
	// houseRober2([]int{1, 2, 3})
	// maxProduct([]int{2, 3, -2, 4})
	maxProduct([]int{-2, 0, -1})

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for _, v := range vals[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func printLinkedList(head *internal.ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// [2,7,11,15], target = 9

func twoSum(nums []int, target int) []int {

	myMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		val, ok := myMap[target-nums[i]]
		if ok {
			a := []int{i, val}
			return a
		} else {
			myMap[nums[i]] = i
		}

	}
	return nil
}

func findMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func findMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func add2List(l1 *ListNode, l2 *ListNode) *ListNode {
	var tempNode *ListNode
	headNode := tempNode
	carry := 0
	for l1 != nil && l2 != nil {
		sum := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		newNode := &ListNode{Val: sum}
		if headNode == nil {
			headNode = newNode
			tempNode = newNode
		} else {
			tempNode.Next = newNode
			tempNode = tempNode.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		newNode := &ListNode{Val: val}
		tempNode.Next = newNode
		tempNode = tempNode.Next
		l1 = l1.Next
	}

	for l2 != nil {
		val := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		newNode := &ListNode{Val: val}
		tempNode.Next = newNode
		tempNode = tempNode.Next
		l2 = l2.Next
	}

	if carry > 0 {
		newNode := &ListNode{Val: carry}
		tempNode.Next = newNode
	}

	return headNode
}

// If repeated, double pointer with a hashset to form a sliding window
func lengthOfLongestSubstring(s string) int {

	p1, p2 := 0, 0
	hashSet := make(map[byte]bool)
	max := 0

	// If hashset exist, if yes, move p1, if not, move p2 and add to hashset
	for p2 < len(s) {
		if _, ok := hashSet[byte(s[p2])]; ok {
			// The goal is to shrink the sliding window, because you've already recorded the max number

			for s[p1] != s[p2] {
				delete(hashSet, s[p1])
				p1++
			}
			delete(hashSet, s[p1])
			p1++
			hashSet[byte(s[p2])] = true

		} else {
			hashSet[byte(s[p2])] = true

			if p2-p1+1 > max {
				max = p2 - p1 + 1
			}
		}
		p2++
	}
	return max
}

/*
1,2,3
3,4,5
1,2,3,3,4,5, even,middleIndex = (arr[len / 2 - 1] + arr[len / 2 ]) /2

1,2,3
3,4,5,6

1 2 3
1 2 3

1 2 3 4
1

4 5 6
1 2 3
nums1, nums2, the median index is always (len(m1) + len(m2)) / 2 if odd,
((len(m1) + len(m2)) / 2 + (len(m1) + len(m2)) / 2 -1) / 2 if even.
so we can set a pointer and the pointer is always less than (len(nums1) + len(nums1)) / 2
and it will for sure stops at (len(nums1) + len(nums1)) / 2

p1, p2 = 0,0, from nums1 and nums2
if nums[]
for i in range (len(nums1) + len(nums1)) / 2



*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)

	if m == 0 {
		return findMedian(nums2)
	}

	if n == 0 {
		return findMedian(nums1)
	}

	p1, p2 := 0, 0
	m1, m2 := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		m2 = m1
		if p1 != m && p2 != n {
			if nums1[p1] < nums2[p2] {
				m1 = nums1[p1]
				p1++
			} else {
				m1 = nums2[p2]
				p2++
			}
		} else if p1 < m {
			m1 = nums1[p1]
			p1++
		} else {
			m1 = nums2[p2]
			p2++
		}
	}

	if (m+n)%2 == 0 {
		return float64(m1+m2) / 2
	} else {
		return float64(m1)
	}

}

func findMedian(nums1 []int) float64 {

	length := len(nums1)
	if length%2 == 0 {
		return (float64(nums1[length/2]) + float64(nums1[length/2-1])) / 2
	} else {
		return float64(nums1[length/2])
	}
}

func mergeSortedArray(nums1 []int, nums2 []int) []int {

	// type PriorityQueue []*internal.Item
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	myPq := make(PriorityQueue, 0)
	for i := 0; i < nums1Len; i++ {
		heap.Push(&myPq, &internal.Item{Value: nums1[i], Priority: nums1[i]})
	}

	for i := 0; i < nums2Len; i++ {
		heap.Push(&myPq, &internal.Item{Value: nums2[i], Priority: nums2[i]})
	}

	mergedArr := make([]int, nums1Len+nums2Len)

	m := myPq.Len()
	fmt.Print("length is %d", m)
	for n := 0; n < m; n++ {
		item := heap.Pop(&myPq).(*internal.Item)
		mergedArr[n] = item.Value
	}
	return mergedArr

}

/* Input: s = "babad"
Output: "bab"
BF:
for(p1 < len, p2<0)
  if(arr[p1] == arr[p2])
*/

func longestPalindromeBF(s string) string {
	sLen := len(s)
	p1, p2 := 0, 0
	for p1 < sLen {
		for p2 = p1 + 1; p2 < sLen; p2++ {
			p3, p4 := p1, p2
			for p3 <= p4 {
				if s[p3] == s[p4] {
					p4--
					p3++
				}
			}
			if p3 > p4 {
				return s[p1:p2]
			}
		}
		p1++
	}
	return s
}

/*
	Input: s = "babad"

Output: "bab"
BF:
for(p1 < len, p2<0)

	if(arr[p1] == arr[p2])

bab
cbbd

p1,p2,p3 = 0
babad
abb
*/
func longestPalindromeFromMiddle(s string) string {
	sLen := len(s)
	p1, p2, p3, maxLen := 0, 0, 0, 0
	var maxStr string

	for p1 < sLen {

		for j := 0; j <= 1; j++ {

			p2, p3 = p1, p1+j

			for p2 >= 0 && p3 < sLen {
				if s[p2] != s[p3] {
					break
				}
				if p3-p2 >= maxLen {
					maxLen = p3 - p2
					maxStr = s[p2 : p3+1]
				}
				p2--
				p3++
			}
		}
		p1++
	}
	return maxStr
}

/*
input: 123
output: 321
123/10 = 12, 3
12/10 = 1, 2,
1/10 = 0, 1

ret = ret*10 + 3, 3
ret = 32
ret = 320 + 1

input:-123
output:-321

123
ret = 3
x = 12

ret = 30 + 2 = 32
x = 1

ret = 320 + 1 = 321
x = 0

345

ret = 345%10 = 5
x = 34

ret = 5*10 + 34%10 = 54
x = 3

ret = 540 + 3



*/

func reverse(x int) int {

	ret := 0

	for x != 0 {
		ret = ret*10 + x%10
		if ret > 2147483647 || ret < -2147483648 {
			return 0
		}
		x = x / 10
	}

	return ret

}

// "432" - 432
// "001" - 1
// "abc12" - 12
/*
result = 0
result = result  * 10 + digit

*/

func myAtoi(s string) int {

	result := 0

	for i := 0; i < len(s); i++ {
		// Skip leading white space
		if s[i] == ' ' {
			continue
		}

		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return result

}

// 123
func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	slice := make([]int, 0)
	for x != 0 {
		ld := x % 10
		slice = append(slice, ld)
		x = x / 10
	}

	len := len(slice)

	m := 0
	n := len - 1

	for m < len && n >= 0 {
		if slice[m] != slice[n] {
			return false
		}
		n--
		m++

	}
	return true

}

// Input stes = ["fla", "flea]
// find the shortest string of the list, becasue the common prefix string cannot be longer than the shortest string
// get the lenght of shortest string
// create a nested loop to loop the entire string list and compare until we reach to the shortest string length
func longestCommonPrefix(strs []string) string {

	myMap := make(map[byte]int)
	shortestStrLen := len(strs[0])
	shortestStr := strs[0]
	sliceEnd := 0
	for i := 0; i < len(strs); i++ {

		if len(strs[i]) < shortestStrLen {
			shortestStrLen = len(strs[i])
			shortestStr = strs[i]
		}
	}
	originalStrLen := len(strs)

	var thisChar byte
	for l := 0; l < shortestStrLen; l++ {

		for i := 0; i < len(strs); i++ {
			thisChar = strs[i][l]
			myMap[thisChar]++
			if myMap[thisChar] == originalStrLen {
				sliceEnd++
				myMap[thisChar] = 0
			}
		}

		if myMap[thisChar] != 0 {
			return shortestStr[0:sliceEnd]
		}

	}
	return shortestStr[0:sliceEnd]
}

/*
input, an array of integers
output, return all the possible triplets array that sum of them is equal to zero
example,
Input: nums = [-1,0,1,2,-1,-4]

Output: [[-1,-1,2],[-1,0,1]]


for i : range len(nums) {
	j = i + 1
	y = target - nums[i]
	for j < len(nums){
		check2Sum(nums[j:])

	}
}

func check2Sum(nums []int) []int {

}

option 2
1. You can sort the the array and use 2 pointers solutions
Input: nums = [-4, -1,-1,0,1,2]
for i := range nums {
	p1, p2:= 0, len(nums) - 1
	off = target - nums[i]

	for p1 < p2 {
		if nums[p1] + nums[p2] + off == target {
			return
		} else if nums[p1] + nums[p2] > off {
			p2 --
		} else {
			p1 ++
		}
	}

}






*/

/*

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
find all sum of subsets is zero

-4 -1 -1 0 1 2
nums[i] + nums[j] + nums[k]

for i := range(nums):
  j = i, k = len(nums) - 1
  for j < k:
	if nums[i] + nums[j] + nums[k] > target {
		k--
	} else if nums[i] + nums[j] + nums[k] < target {
		j++
	} else {
		append to results
	}

*/

func threeSum(nums []int) [][]int {
	var j, k int
	var results = [][]int{}
	sort.Ints(nums)

	for i := range nums {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j = i + 1
		k = len(nums) - 1

		for j < k {

			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				results = append(results, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}

				for k > j && nums[k] == nums[j-1] {
					k--
				}
				k--
				j++
			}
		}
	}
	return results
}

/*




 */

func threeSumBC(nums []int) [][]int {
	res := [][]int{}
	var dfs func(nums []int, start int, temp []int, target int)
	sort.Ints(nums)
	dfs = func(nums []int, start int, temp []int, target int) {

		if len(temp) > 3 {
			return
		}

		if target == 0 && len(temp) == 3 {
			tempArr := make([]int, len(temp))
			copy(tempArr, temp)
			res = append(res, tempArr)

		}

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(nums, i+1, temp, target-nums[i])
			temp = temp[:len(temp)-1]
		}

	}

	dfs(nums, 0, []int{}, 0)

	return res

}

/*
Remove nth node from the end of list
example 1, n0 -> n1 -> n2, position = 1, output: n0->n1
example 1, n0 -> n1 -> n2, position = 2, output: n0->n2
example 1, n0 -> n1 -> n2, position = 0, output: n0->n1
example , n0 -> n1 -> n2 ->    n3 ->                  n4, position = 2, output: n0->n1->n3->n4
                p1    p1.next  p1.next.next           p2

appraoch: use two pointers, with position n distance, walk through the list unitl the 2nd pointer.next is null.
1. init 2 pointers, let p1.next and p2 has position n distance
      n0 -> n1 -> n2 -> n3 -> n4, n = 2
	p1.next        p2
2. walk through the list, until p2.next is null
	    n0 -> n1 -> n2 -> n3 -> n4,
		       p1      p1.next    p2

3. do the skip
 p1.next = p1.next.next


*/

func removeNthNodeFromEndofList(head *ListNode, n int) *ListNode {

	p1, p2 := head, head

	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	if p2 == nil {
		return head.Next
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next

	return head

	// p1 := &ListNode{Val: 0, Next: head}
	// p2 := head
	// i := n
	// for i > 0 {
	// 	p2 = p2.Next
	// 	i--
	// }

	// for p2.Next != nil {
	// 	p1 = p1.Next
	// 	p2 = p2.Next
	// }

	// rtNode := &ListNode{Val: p1.Next.Val, Next: p1.Next}
	// p1.Next = p1.Next.Next
	// return rtNode

	// slow, fast := head, head
	// for i := 0; i < n; i++ {
	//     fast = fast.Next
	// }
	// if fast == nil {
	//     return head.Next
	// }
	// for fast.Next != nil {
	//     fast = fast.Next
	//     slow = slow.Next
	// }
	// slow.Next = slow.Next.Next
	// return head
}

/*
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// (), valid
// {}[],valid
// {[]}, valid
// {[}, invalid

Requirements:
only 3 sets of characters, so we can hard code the check.

Approach, use stack.
input: "{[]}", s -> empty
if(open bracbracketses) then stack.push(close brackets)
else stack.pop and check if they are same,
input: "[]}", s -> "}"
input: "]}", s -> "}]"
"){"
")(){}"
"[])"
*/

// TODO
func isValidParenthese(s string) bool {

	inputLenth := len(s)
	myLocalStack := Stack{}

	if inputLenth == 0 {
		return true
	}

	for i := 0; i < inputLenth; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			myLocalStack.Push(s[i])
		} else {
			if myLocalStack.Size() != 0 && (myLocalStack.Peek() == '(' && s[i] == ')' ||
				myLocalStack.Peek() == '{' && s[i] == '}' ||
				myLocalStack.Peek() == '[' && s[i] == ']') {
				myLocalStack.Pop()
			} else {
				return false
			}
		}

	}

	if myLocalStack.Size() != 0 {
		return false
	} else {
		return true
	}

}

// "(){}"
// "){}"
// "({}"
func isValidParenthes2(s string) bool {

	inputLenth := len(s)
	myLocalStack := Stack{}

	for i := 0; i < inputLenth; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			myLocalStack.Push(s[i])
		} else if (s[i] == ')' && myLocalStack.Peek().(rune) == '(') || (s[i] == ']' && myLocalStack.Peek().(rune) == '(') || (s[i] == '}' && myLocalStack.Peek().(rune) == '{') {
			myLocalStack.Pop()
		} else {
			return false
		}

	}

	if myLocalStack.Size() > 0 {
		return false
	}
	return true

}

/*
Given 2 sorted list node:
1->2->4
3->5->7
output: 1,2,3,4,5,7
use 2 pointers each point to the list, and then go thourgh list until one of the pointer reach to null

for(p1!=null && p2!=null)
	if(p1.value <= p2.value)
		newHead = Node(p1.value)
		p1 = p1.next
	else
		newHead = Node(p2.value)
		p2 = p2.next
	newHead = newHead.next

if(p1 != null)


if(p2 != null)

return newHead

*/

func mergedSortedList(list1 *ListNode, list2 *ListNode) *ListNode {

	newListNode := &ListNode{}
	rtHead := newListNode

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newListNode.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next

		} else {
			newListNode.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		newListNode = newListNode.Next

	}

	for list1 != nil {
		newListNode.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		newListNode = newListNode.Next
	}

	for list2 != nil {
		newListNode.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		newListNode = newListNode.Next
	}

	return rtHead.Next

}

/*
1->2->4
3->5->7
8->9->10

User merge sort to divide the problem to merge only 2 lists or 1
Merge sort:
 1. Recursion, likely a helper function
 2. condition when to call merge2List
 3. condition to return if there is only 1 list
 4. find mid, start and end pointer
*/
func mergeKLists(originalList []*ListNode) *ListNode {
	if originalList == nil {
		return nil
	}

	return mergeKSortedListHelper(originalList, 0, len(originalList)-1)

}

// where recursion happen
func mergeKSortedListHelper(list []*ListNode, start int, end int) *ListNode {

	if start == end {
		return list[start]
	}

	if start == end+1 {
		return mergedSortedList(list[start], list[end])
	}

	// keep dividing

	mid := start + (end-start)/2

	leftList := mergeKSortedListHelper(list, start, mid)
	rightList := mergeKSortedListHelper(list, mid+1, end)

	return mergedSortedList(leftList, rightList)

}

/*
Another approach is to jump 2 items
*/
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	for len(lists) > 1 {

		var mergedList []*ListNode
		for i := 0; i < len(lists); i += 2 {
			m := i
			n := i + 1

			var aa *ListNode

			if n < len(lists) {

				aa = mergedSortedList(lists[m], lists[n])

			} else {
				aa = mergedSortedList(lists[m], nil)
			}

			mergedList = append(mergedList, aa)
		}

		lists = mergedList

	}

	return lists[0]
}

/*
input list = 1 -> 2 -> 3 -> 4
output = [2,1,4,3]
Constraint: cannot modify the valie, but switch the node.
*/

// Version that directly change value
func SwapPair(list *ListNode) *ListNode {
	var temp int

	var p1, p2 *ListNode
	p1 = list
	p2 = list.Next

	for p2 != nil {
		temp = p1.Val
		p1.Val = p2.Val
		p2.Val = temp
		p1 = p1.Next
		p2 = p2.Next
	}

	return list

}

/*Version that not change value but change nodes
input  = 1 -> 2 -> 3 -> 4
output = 2 -> 1 -> 4 -> 3
*/

func SwapPairNode(list *ListNode) *ListNode {

	p1 := list
	p2 := p1.Next

	p3 := p1.Next.Next

	for p3 != nil {
		p2.Next = p1
		p1.Next = p3.Next
		p1 = p3
		p2 = p3.Next

	}

	return nil

}

/*
Remove duplicated integer in an array,
Return the number of unique items.
example:
[1,1,2] -> [1,2,_], return 2

Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.


solution 1:
use 2 pointers, p1 and p2
p1 = 0
p2 = p1 + 1
for(p2 < len(nums))
	if(p2 == p1)
	  p2++
	else
	  p1++
	  p1 = p2
	  p2 ++
for(p1 < len(nums))
  reset(p1)




*/

func removeDuplicatedArr(nums []int) int {

	p1 := 0
	p2 := 1

	for p2 < len(nums) {
		if nums[p1] != nums[p2] {
			p1++
			nums[p1] = nums[p2]
		}
		p2++
	}
	p1++
	rt := p1

	for p1 < len(nums) {
		nums[p1] = '_'
		p1++
	}

	return rt

}

/*
arr = [2,3,3,2], val = 3
arr1 = [2,2,-,-]
[0,1,2,2,3,0,4,2], 2
*/

func removeElement(nums []int, val int) int {

	readIndex, writeIndex := 0, 0

	for readIndex < len(nums) {
		if nums[readIndex] != val {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
		readIndex++
	}

	rt := len(nums) - writeIndex

	for writeIndex < len(nums) {
		nums[writeIndex] = '_'
		writeIndex++
	}

	return rt

}

/*
find the first occurance of sub string
Input: haystack = "sadbutsad", needle = "but", return 3 as index.
haystack="aaa", needle="aaaa"
"mississippi", "issipi"
*/

func Strstr(haystack string, needle string) int {
	neddleLen := len(needle)
	haystackLen := len(haystack)

	if neddleLen > haystackLen {
		return -1
	}

	for i := 0; i < haystackLen; i++ {
		temp := i

		if temp > haystackLen-neddleLen {
			return -1
		}

		for j := 0; j < neddleLen; j++ {
			if haystack[temp] != needle[j] {
				break
			}
			temp++
		}
		if temp-i == neddleLen {
			return i
		}
	}

	return -1

}

/*
binary serach with sorted array
*/
func binarySearchRecrusion(nums []int, target int) int {

	return binarySearchHelper(nums, 0, len(nums), target)
}

func binarySearchHelper(nums []int, low int, high int, target int) int {

	mid := low + (high-low)/2
	if nums[mid] == target {
		return mid
	}

	if target >= nums[mid] {

		return binarySearchHelper(nums, mid, high, target)

	}

	if target < nums[mid] {
		return binarySearchHelper(nums, low, mid-1, target)
	}

	return -1

}

/*
Given a sorted array with rotation and target, return index
[4,5,6,7,0,1,2],target = 0, return 4

low, mid, high

	for low <= high {
	   if nums[mid] == target:
	      found

	    if nums[low] <= nums[mid] {
	       if target <= nums[mid] and target >  nums[low], normal binary search, apply binary search
		      high = mid
			  mid = low + (high-low) / 2
		   else
		      low = mid

	   //target is in the right
		else:
	       if target > nums[mid] and target <  nums[high]
		      low = mid
			  mid = low + (high-low) / 2
		   else
		      high = mid

}
*/
func searchRotatedArray(nums []int, target int) int {
	return 1
}

/*

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

*/

func searchRange(nums []int, target int) []int {

	first, last := -1, -1

	if len(nums) == 0 {
		return []int{first, last}
	}

	for i, _ := range nums {

		if target != nums[i] {
			continue
		}

		if first == -1 {
			first = i
		}

		last = i

	}

	return []int{first, last}

}

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

0,1,2,3,5  4
*/

func searchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}
	for i, _ := range nums {

		if target == nums[i] || target < nums[i] {
			return i
		}

	}

	return len(nums)

}

/*
[1,3,5,6], target = 5
l = 2, r = 3

[1,3,5,6], target = 7
l = 2, r = 3, mid = 2

[1,3,5,6], target = 0
*/
func searchInsertBin(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l < r {

		mid := l + (r-l)/2
		if target < nums[mid] {
			r = mid - 1

		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}

	}
	if nums[l] < target {
		return l + 1
	}
	return l

}

/*

Given an integer array nums, find the
subarray with the largest sum, and return its sum.
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
tip: only positive number can contirbute to max number, so we can ignore negative number

*/

func maxSubArray(nums []int) int {

	current_max := 0
	global_max := math.MinInt
	for i, _ := range nums {
		current_max = current_max + nums[i]
		if current_max <= 0 {
			current_max = 0
		} else {
			global_max = checkBig(global_max, current_max)
		}

	}
	return global_max

}

func checkBig(a int, b int) int {

	if a > b {
		return a
	} else {
		return b
	}

}

/*
Given a tree, find the lowest common ancestor.


*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// find base cases, node cannot boe lower than its ancesotr, ancestor only is node it self
	// or above it
	if p == root || q == root {
		return root
	}
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)

	if leftNode != nil && rightNode != nil {
		return root
	}

	return nil

}

/*
Rotate linkedList
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]
Set 2 pointers with K gap
e.g, p1 is 3, and p2 is 5,
p2.next = head
rt = p1.next
p1.next = nil
return rt
*/
func rotateRight(head *ListNode, k int) *ListNode {

	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p2 = p2.Next

	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = head
	rt := p1.Next
	p1.Next = nil
	return rt

}

/*
Sort 3 colors.
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
*/
func sortColors(nums []int) {
	arr := [3]int{0}
	for i, _ := range nums {
		arr[nums[i]]++
	}

	var i int = 0
	// var j int = 0
	for i = 0; i < arr[0]; i++ {
		nums[i] = 0
	}

	for j := 0; j < arr[1]; j++ {
		nums[i] = 1
		i++
	}

	for i < len(nums) {
		nums[i] = 2
		i++
	}

}
func inorderTraversal(root *TreeNode) []int {

	arr := []int{}
	inorderTraversalHelper(root, &arr)
	return arr

}

func inorderTraversalHelper(root *TreeNode, arr *[]int) {

	if root == nil {
		return
	}

	inorderTraversalHelper(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorderTraversalHelper(root.Right, arr)
}

/*
Valid means if in-order is ascending order.
If root.Val is greater than left but less than right

*/

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left != nil && root.Val <= root.Left.Val {

		return false

	}

	if root.Right != nil && root.Val > root.Right.Val {
		return false

	}
	val1 := isValidBST(root.Left)

	val2 := isValidBST(root.Right)

	return val1 || val2

}

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	q := []*TreeNode{root.Left, root.Right}

	for len(q) != 0 {
		left := q[0]
		right := q[1]
		q = q[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		q = append(q, left.Left)
		q = append(q, right.Right)

		q = append(q, left.Right)
		q = append(q, right.Left)

	}
	return true
}

// func checkSymeetric(arr []*TreeNode) bool {
// 	i := 0
// 	j := len(arr) - 1

// 	if len(arr) == 1 {
// 		return true
// 	}

// 	if len(arr)%2 != 0 {
// 		return false
// 	}

// 	for i <= j {
// 		if arr[i] != arr[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true

// }

/*
Level order tree,
BST.
DS needed, queue, a slice.
q.enqueu(root)

	for !q.empty(){
	  new slice
	  size = q.size()
	  node = q.dequeu()
	  for i = 0, i< size, i++
	  	node = q.poll()
		slice.append(node.val)
	  	q.enqueue(root.right)

	  checkSymeetric(slice)

}
*/
func levelOrder(root *TreeNode) [][]int {
	// q := internal.InitQueue(100)

	q := []*TreeNode{root}

	levelSlices := [][]int{}
	if root == nil {
		return levelSlices
	}

	for len(q) != 0 {
		levelSize := len(q)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		levelSlices = append(levelSlices, level)
	}
	return levelSlices
}

/*
1. A method to add linkedList at end
2. A method to pre-order travesal
3. Remove Left node
4. Use another variable to store the list
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var myList *TreeNode

	flattenHelper(root, myList)

}

func flattenHelper(root *TreeNode, myList *TreeNode) {

	if root == nil {
		return
	}
	myList = addListTail(myList, root.Val)
	flattenHelper(root.Left, myList)
	flattenHelper(root.Right, myList)

}

func addListTail(head *TreeNode, val int) *TreeNode {

	if head == nil {
		return &internal.TreeNode{Val: val}
	}
	head.Right = addListTail(head.Right, val)
	return head

}

/*
Find max depth of the tree
DFS search
Post-order serach
*/
func maxDepth(root *TreeNode) int {
	if root != nil {
		return maxDepthHelper(root)
	}
	return 0

}

func maxDepthHelper(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := maxDepthHelper(root.Left) + 1
	right := maxDepthHelper(root.Right) + 1
	return checkMax(left, right)

}

func checkMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
Invert tree
*/
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		temp := root.Right
		root.Right = root.Left
		root.Left = temp
		invertTree(root.Left)
		invertTree(root.Right)
	}

	return root
}

/*
[1,1,1,2,2,3], 2

	m = {
		1 -> 3,
		2 - >2,
		3 -> 1
		-1: 2
	}

[_,1,2,3,_]

[[],[3], [2,-1],[1]]
*/
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, val := range nums {
		m[val]++
	}

	bucket := make([][]int, len(nums)+1)

	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}

	ans := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, val := range bucket[i] {
			if k > 0 {
				ans = append(ans, val)
				k--
			}
		}
	}

	return ans
}

func hasCycleHashTable(head *ListNode) bool {

	m := make(map[*ListNode]bool)

	currentNode := head

	for currentNode != nil {
		if m[currentNode] == true {
			return true
		}
		m[currentNode] = true
		currentNode = currentNode.Next
	}
	return false

}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumNumbersHelper(root, 0)

}

/*
1
2 3

1

	2
*/
func sumNumbersHelper(root *TreeNode, sum int) int {

	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}

	sum = sum*10 + root.Val
	left := sumNumbersHelper(root.Left, sum)
	right := sumNumbersHelper(root.Right, sum)

	return left + right

}

func lowestCommonAncestorp(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if p == root || q == root {
		return root
	}
	left := lowestCommonAncestorp(root.Left, p, q)
	right := lowestCommonAncestorp(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil

}

/*
if p == root || q == root
 return  root

 if root.val>p.val && root.val < q.val || root.val > q.val && root.val < p.val
  return root





*/

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	if p == root || q == root {
		return root
	}

	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestorBST(root.Right, p, q)
		if right != nil {
			return right
		}
	}

	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestorBST(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	return root

}

/*

Input: s = "anagram", t = "nagaram"
a->3
n->1
g->1
r->1
m->1

lettersAllLen = len(m) = 5
for i in range t
  if map[t[i]] not exist ; return false

  map[t[i]] --

  if map[t[i]]  == 0 {
	lettersAllLen--
  }

*/

/*
Input: strs = ["eat","tea","tan","ate","nat","bat"]
*/
func groupAnagrams(strs []string) [][]string {

	set := make(map[[26]int][]string)

	for _, val := range strs {
		var chars [26]int
		for _, c := range val {
			chars[c-'a']++
		}
		set[chars] = append(set[chars], val)
	}

	var result [][]string

	for _, v := range set {
		result = append(result, v)
	}

	return result
}

/*
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

for i in range nums
BF:

	  j = 0
	  product = 1
	  while j < len(nums){
		if j != i {
			product = product * nums[j]
		}
	  }

Optimized:

nums[0] = nums[1] * nums[2] * nums[3] = y1
nums[1] = nums[0] * nums[2] * nums[3], y1 / nums[1] * nums[0] = y2
nums[2] = nums[0] * nums[1] * nums[3], y2 / nums[2] * nums[1] = y3
nums[3] = nums[0] * nums[1] * nums[2], y3 / nums[3] * nums[2] = y4
1,2,3,4
*/
func productExceptSelf(nums []int) []int {

	product := 1
	productArr := make([]int, len(nums))
	for _, val := range nums[1:] {
		product = product * val
	}
	productArr[0] = product

	for i := 1; i < len(nums); i++ {
		productArr[i] = product / nums[i] * nums[i-1]
		product = productArr[i]
	}
	return productArr

}

/*

Input: nums = [2,20,4,10,3,4,5]
2,3,4,5

Output: 4

if (nums[i] - 1) not exist in map



*/

func longestConsecutiveBF(nums []int) int {

	m := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	maxSequenceLen := 0
	for i := 0; i < len(nums); i++ {

		val := m[nums[i]-1]
		if val == false {
			// nums[i] is the starting sequence
			seq := 1
			for {
				if _, ok := m[i+seq]; ok {
					seq++
					continue
				}

				maxSequenceLen = max(maxSequenceLen, seq)
				break
			}

		}

	}

	return maxSequenceLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getConcatenation(nums []int) []int {

	ans := make([]int, 2*len(nums), 2*cap(nums))

	for i, val := range nums {
		ans[i] = val
		ans[i+len(nums)] = val
	}

	return ans

}

/*
Input: s = "racecar", t = "carrace"

Output: true
r 2
a 2
c 2
e 1


*/

func validAngram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)

	for _, item := range s {
		m[item]++
	}

	for _, item := range t {
		if _, ok := m[item]; ok {
			m[item]--
			if m[item] == 0 {
				delete(m, item)
			}
		} else {
			return false
		}
	}

	if len(m) == 0 {
		return true
	} else {
		return false
	}
}

/*
Input:
nums = [3,4,5,6], target = 7

if m[target - nums[i]] is nill

	m[i] = i

else

	return i, m[target - nums[i]]

Output: [0,1]
*/
func twoSum1(arr []int, target int) []int {

	myMap := make(map[int]int)

	for index, val := range arr {
		offset := target - val

		if val, ok := myMap[offset]; ok {
			return []int{val, index}
		} else {
			myMap[val] = index
		}

	}
	return nil

}

/*
Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

 0-25
 a, c, t
 1  1  1

 Create a map, key is the alphabet array with character frequen for each work
 e.g, act, cat will be same key and group together.





*/

func groupAnagramsDetail(strs []string) [][]string {

	m := make(map[[26]int][]string)

	for _, val := range strs {
		var alphabetArr [26]int

		for _, item := range val {
			alphabetArr[item-'a']++
		}
		m[alphabetArr] = append(m[alphabetArr], val)
	}

	var result [][]string

	for _, val := range m {
		result = append(result, val)
	}

	return result

}

/*
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]

1,2    3,4

Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	newArr := make([]int, m+n)

	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			newArr[k] = nums1[i]
			i++
		} else {
			newArr[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		newArr[k] = nums1[i]
		i++
	}

	for j < n {
		newArr[k] = nums2[j]
		j++
	}

}

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] where nums[i] + nums[j] + nums[k] == 0, and the indices i, j and k are all distinct.

The output should not contain any duplicate triplets. You may return the output and the triplets in any order.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]

Output: [[-1,-1,2],[-1,0,1]]


sorted
[-4,-1,-1,0,1,2,2]

[1,2,3] 6

result = left + right
target_off = target - nums[i]

if result < target_off,  left ++
if result > target_off, right --
if result == target_off

Using a set(map)

for i in range(0-len(nums)-1)
	for j in range(i+1-len(nums)-1)
       off = target - nums[i] - nums[j]
	   if off in hashset:
	       return true

		hashset.add(nums[j])

*/

// func threeSum(nums []int) [][]int {

// }

/*

if s[j] != alphanumeric
 j--
 continue

if s[i] != alphanumeric
 i++
 continue

if s[j != s[i]
 return false


 return true


*/

func isPalindromestr(s string) bool {
	i, j := 0, len(s)-1

	for i != j {
		if !isLetter(s[i]) {
			i++
			continue
		}
		if !isLetter(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}

	}

	return true

}

func isLetter(char byte) bool {
	if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
		return false
	} else {
		return true
	}
}

/*
input 2,1,1,0,0,1
      1,1,1,0,0,2
input 0,0,2,2,1

arr []int
0,1,2
3 2 1


i, left, right

for i < right
  if arr[i] == 0
  swap(i,left)
  i++, left ++

  if arr[i] == 1
   i++



  if arr[i] == 2
  swap(i,right)
  right --


*/

func DutchNationalFalgSwap(arr []int) []int {

	i, left, right := 0, 0, len(arr)-1

	for i < right {
		if arr[i] == 0 {
			temp := arr[i]
			arr[i] = arr[left]
			arr[left] = temp
			i++
			left++

		} else if arr[i] == 1 {
			i++

		} else {
			temp := arr[i]
			arr[i] = arr[right]
			arr[right] = temp
			right--

		}

	}
	return arr

}

/*
Input: nums = [2,20,4,10,3,4,5]

Output: 4, 2,3,4,5

Brutal Force:

result = [][]
for i in range(0-len-1)

	    temp = 1
		temparr={nums[i],...}
		for j in range(1-len-1)
		   if num[j] == nums[i] + temp:
		      append(temparr, num[j])
			  temp ++

		result.append(temparr)

Optimized:
Pre-load the data, as we only care about the starging sequence
Like 2sum question

for i n range nums{

	// It's the starting sequence
	if set[nums[i] - 1 ] != true {
		seq = 1
		for i < len(nums){
			if set[nums[i] + seq] == true {
				max = Match.max(small, seq+1)
			}
			seq++
			i++
		}


	}

}
*/
func longestConsecutiveBF2(nums []int) int {
	numSet := make(map[int]bool)

	for _, val := range nums {
		numSet[val] = true
	}

	for _, val := range nums {
		if numSet[val-1] != true {

		}
	}
	return 1
}

/*
Given an array of integers numbers that is sorted in non-decreasing order.

Return the indices (1-indexed) of two numbers, [index1, index2], such that they add up to a given target number target and index1 < index2. Note that index1 and index2 cannot be equal, therefore you may not use the same element twice.

There will always be exactly one valid solution.

Your solution must use
O
(
1
)
O(1) additional space.

Example 1:

Input: numbers = [1,2,3,4], target = 3

Output: [1,2]
Explanation:
The sum of 1 and 2 is 3. Since we are assuming a 1-indexed array, index1 = 1, index2 = 2. We return [1, 2].

Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

1 0
2 1
3 2
4 3

because it' sorted array, so also use 2 pointers,left and right
left =0. right = n-1

add = left + right
if add < target, left++
if add > target, right ++
if add == target

	for left < right{
		add = left + right
		if add < target, left++

if add > target, right ++
if add == target

}
Input: numbers = [2,3,4], target = 5
Input: numbers = [1,2,3,4], target = 3
*/
func twoSumArry(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		result := numbers[left] + numbers[right]
		if result < target {
			left++
		} else if result > target {
			right--
		} else {
			return []int{left, right}
		}
	}
	return []int{}
}

/*
You are given an integer array heights where heights[i] represents the height of the
i
t
h
i
th

	bar.

You may choose any two bars to form a container. Return the maximum amount of water a container can store.
Input: height = [1,7,2,5,4,7,3,6]

Output: 36
Example 2:

Input: height = [2,2,2]

Output: 4

BF:
max = verysamllnumber
for i -> (0,n-1)

	for j -> (i+1,n-1)
		x = j -1
		y = math.min(i,j)
		max = math.max(max, x*y)

Optimized:
[1,5,3,4]
two pinters, left and right
result = min(nums[left], nums[right]) * (right - left)
if nums[left] < nums[right] left ++
if nums[left] >= nums[right] right --
*/
func maxWater(heights []int) int {

	left, right := 0, len(heights)-1
	max := math.MinInt

	for left < right {
		result := findMin(heights[left], heights[right]) * (right - left)
		max = findMax(max, result)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}

	}
	return max

}

/*
Input: prices = [10,1,5,6,7,1]
Input: prices = [10,3,5,1,7,12]

Output: 6, nums[1], nums[4]
i = 0, j =1
if num[j] - nums[i] < 0, i++,j++
else j++, max = findMax(max, nums[j]-nums[i])

[2,1,2,1,0,1,2], 2
*/
func findMaxProfi(prices []int) int {
	i, j := 0, 1
	max := 0

	for j < len(prices) {
		if prices[j]-prices[i] < 0 {
			i++
			continue
		}
		max = findMax(max, prices[j]-prices[i])
		j++
	}

	return max

}

/*
Input: s = "zxyzxyz"

Output: 3

"abca"
3

"aba"
2

use a hashset

	if !set.contains(s[i]){
		set.add(s[i])
		count++
		max = findMax(max,count)
	} else {

	clearn hash set
	set.add(s[i])
	count = 1

}
"dvdf"
"ababc"
slow, fast = 0, 1

	for fast < n {
		if nums[fast] == nums[slow]{
			slow++
			fast = slow + 1
		}
		max = findMax(max, fast - slow+1)
		fast ++
	}

	abcabcbb
	ababc
	dvdf

	s = 0, f = 1, max = 1
	s = 1, f =2,

pwwkew

		func lengthOfLongestSubstring(s string) int {
	    curBest := 0
	    set := map[byte]bool{}
	    l, r := 0, 0
	    for r < len(s) {
	        if _, ok := set[s[r]]; !ok { // can grow window
	            set[s[r]] = true
	            r++
	            curBest = max(curBest, r - l)
	        }else{
	            delete(set, s[l])
	            l++
	        }
	    }
	    return curBest
	}
*/
func lengthOfLongestSubstringWithoutRepeatingChar(s string) int {

	if len(s) == 0 {
		return 0
	}
	set := make(map[byte]bool)
	slow, fast, max := 0, 0, 1
	for fast < len(s) {
		if _, ok := set[s[fast]]; ok {
			delete(set, s[fast])
			slow++
			// fast = slow
			// set = make(map[byte]bool)

		} else {
			set[s[fast]] = true
			fast++
			max = findMax(max, fast-slow)

		}
	}
	return max
}

/*

You are given a string s consisting of only uppercase english characters and an integer k. You can choose up to k characters of the string and replace them with any other uppercase English character.

After performing at most k replacements, return the length of the longest substring which contains only one distinct character.

Example 1:

Input: s = "XYYX", k = 2

Output: 4

left=0, right=0

xxyyx

for fast < n{
	temp = k
	if s[left] == s[right]{
		max = findMax(max, right-left+1)
		right++

	} else {
		if temp > 0{
			max = findMax(max, right-left+1)
			right ++
			temp--

		} else {
			left ++
			right = slow
		}

	}
}
ABAB, 2
AABABBA,1

AAAB


ABBB, 2
*/

func longestRepeatingChar(s string, k int) int {
	left, right := 0, 0
	max := 1

	temp := k
	for right < len(s) {

		if s[left] == s[right] {
			max = findMax(max, right-left+1)
			right++
		} else if temp > 0 && s[left] != s[right] {
			// change left first
			if temp == k {

			}
			max = findMax(max, right-left+1)
			temp--
			right++

		} else {
			left++
			right = left
			temp = k
		}
	}
	return max

}

/*
You are given two strings s1 and s2.

Return true if s2 contains a permutation of s1, or false otherwise. That means if a permutation of s1 exists as a substring of s2, then return true.

Both strings only contain lowercase letters.

Example 1:

Input: s1 = "abc", s2 = "lecabee", "eaecabee"

BF:
sort both s1 and s2
abc, eeeabcppp

	for i -> range len(s2) -1{
		for j in range len(s1){
			if s2[i] != s1[j]{
				break
			} else {
				k++
			}
		}

}

Input: s1 = "abc", s2 = "lecabee", "eaecabee"

make a windows size equal to s1, 3
left-0, right = 3-1 = 2

	for right < len(s1){
		for i = left, i < =right{
			temp := 1
			if !map.contains(s2[]) {
				//move window,
				left ++
				right++
			} else {

				temp ++
			}

			if temp == len(s1) return true
		}

}
Input: s1 = "abc", s2 = "lecabee", "eaecabee"
abc, aaaaaa
hello,ooolleoooleh
ab, cab
*/

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, len(s1)-1

	set := make(map[byte]int)

	for _, val := range s1 {
		set[byte(val)]++
	}
	for right <= len(s2)-1 {
		temp := 0
		copiedMap := make(map[byte]int)
		for key, value := range set {
			copiedMap[key] = value
		}
		for i := left; i <= right; i++ {
			if _, ok := copiedMap[byte(s2[i])]; ok {
				temp++
				copiedMap[byte(s2[i])]--
				if copiedMap[byte(s2[i])] == 0 {
					delete(copiedMap, s2[i])
				}
			} else {
				break
			}
		}
		if temp == len(s1) && len(copiedMap) == 0 {
			return true
		}
		left++
		right++
	}

	return false

}

/*

Given two strings s and t, return the shortest substring of s such that every character in t, including duplicates, is present in the substring. If such a substring does not exist, return an empty string "".

You may assume that the correct output is always unique.

Example 1:

Input: s = "OUZODYXAZV", t = "XYZ"

Output: "YXAZ"


BF:
for i =0, i < n-1 i++ {
	for j = i, j< n-1, j++ {
		max = findMax(max, checkIftInS(s[i:j], t))
	}
}
// check if char in s2 are all in s1
func checkIftInS(s1, s2){
	creaet map, s2[2]->freqency
	for range s1{
		if map[s1[i]]{
			map[s1[i] --
		}
		if map[s1[i]] == 0 {
			delete
		}
	}
	if map is emoty
	return len(s1)
	else{
		return 0
	}
}

sliding window

tMap
{
	X:1
	Y:1
	Z:1
}





Input: s = "OUZODYXAZV", t = "XYZ"

Output: "YXAZ"
sMap
left = 0, right  = 0
for right < n-1{
	if t[right] == ok {
		sMap[val]++
		if s[:i] has characert of tMap{
			get Max len
			while(left < right){
				if sMap[left] is ok{
					sMap[left]--
					break
				}
				left ++
			}
		}
	}


}


*/

// func minWindow(s string, t string) string {

// }

/************
Tree
pre-order, in-oreder, post-oreder
Traverse using itterations.
1
/\
2 3

pre-order, 1,2,3, root, left, right
in-order, 2,1,3, left, then root, then right
post-oder, 2,3,1, left, then right, then root

Pre order:
Using itterations basically needs a DS to store some state.
stack.push(root)
while(stack is not empty) {
	r = stack.pop()
	print(r)
	stack.push(r.right)
	stack.push(r.left)
}

In-order:
traverse left,
visit node
traverse right

currentnode = root
stack.push(currentNode)
currentNode = currentNode.left
stack.push(currentNode)
currentNode = currentNode.left

currentNode == nill
newNode = stack.pop()
print(newNode)











**************/

/*
You are given the head of a singly linked-list.

The positions of a linked list of length = 7 for example, can intially be represented as:

[0, 1, 2, 3, 4, 5, 6]
0 6 1 5 2 4 3

0123
0312
1. find the middle pointer,
2. reverse 2nd half of the list
3. merge first half and 2nd half of list

1 2 3 4
5 6 7 8
l1_temp 2

l1_temp = l1.next
l1.next = l2
l1 = l1_temp
l2_temp = l2.next
l2.next = l1
l2 = l2_temp
*/
func reOrderList(head *ListNode) {

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

}

func mergeTwoList(l1 *ListNode, l2 *ListNode) {

	for l1 != nil && l2 != nil {
		l1_temp := l1.Next
		l1.Next = l2
		l1 = l1_temp
		l2_temp := l2.Next
		l2.Next = l1_temp
		l2 = l2_temp
	}

}

func reverseLinkedList(head *ListNode) *ListNode {

	cur := head
	var pre *ListNode

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre

}

/*

Input: nums =
[3,4,5,6,1,2]
[5,6,1,2,3,4]
[6 1 2 3 4 5]
[4 5 6 1 2 3]
[3 4 5 6 1 2]
[5 6 1 2 3 4]
[1,2,3,4,5,6]
1 2 3 4 5 6
6 1 2 3 4 5
2 3 4 5 6 1
Output: 1

for left < right{

}
if right > left
  getLeft

if mid < right , it's on the left part of the array

if mid > left, it's on the right part of array


[3,4,5,6,1,2]
[5,6,1,2,3,4]
[6 1 2 3 4 5]
[1,2,3,4,5,6]

0 1 2 3
[4 1 2 3]
1 2 3 0

4 1 2, r = 2, left = 0, mid =1

mid = 1, left = 0, right =3,
left = 0 , right = 0, mid = 0

*/

func findMinRotaedSortArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[right] > nums[left] {
			return nums[left]
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]

}

/*


1 2 3 4
1 2 3 4 5
1 2 3 4 5 6

binarySerachHelper(arr, left, right, target) {

	mid = left + (right-left) / 2
	if(arr[mid] == target) return mid

	if target >= arr[mid] {
		return binarySerachHelper(arr, mid+1, right)

	}

	if target < arr[mid]{
		return binarySerachHelper(arr, left, mid1)
	}
	return nil


}
Input: nums = [3,4,5,6,1,2]

left = 0, right = 5,  mid = 2 -> left = 3, mid = 4


Input: nums = [5,6,1,2,3,4]
searchInRotaedSortArray(arr []int, target int) {


	mid = right / 2
	for left < right {

		if nums[left] < nums[mid]{
			if target <= mid{ // normal binary serach
				binarySerachHelper(arr, left, mid-1)
		}  else {
				left = mid -1 1
			}
		}else  {
			if  target >= mid) // normal binary serach
			{
				binarySerachHelper(arr, mid + 1, right)
			} else {
				right = mid + 1

			}
		}


	}

}

[5,6,1,2,3,4]
[2,3,4,5,6,1]

*/

func searchInRotaedSortArray(arr []int, target int) int {

	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[left] <= arr[mid] {

			if target >= arr[left] && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			if target >= arr[mid] && target < arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1

}

func invertBinaryTree(node *TreeNode) {
	if node == nil {
		return
	}

	temp := node.Left.Val
	node.Left.Val = node.Right.Val
	node.Right.Val = temp

	invertBinaryTree(node.Left)
	invertBinaryTree(node.Right)

}

func maxDepth2(node *TreeNode) int {
	return maxDepth2Helper(node, 0)
}

func maxDepth2Helper(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	leftNode := maxDepth2(node.Left) + 1
	rightNode := maxDepth2(node.Right) + 1

	return findMax(leftNode, rightNode)

}

/*
Find the max sum for any path.
Find edge case
find out which depth is big
if node == nil {
	return 0
}
left = maxPathSum(node.left)
righ = maxPathSum(node.right)

if left =< 0 {
	left = 0
}


leftMaxDepth = left + node.Val
rightMaxDepth = right + node.Val
return findMax(leftMaxDepth,rightMaxDepth)




*/

func maxPathSum(node *TreeNode) int {
	return -1
}

/*
Invert binary tree.
The goal is to invert entire node inlcuding value

1
2 3

    1
 3    2
4 5   6 7

1. Is this a Top-Down or Bottom up question.
It's top down, so use pre-order traverse.
define base case, node is nil

if node == nill; return nil

if node.left && node.right
temp:=node.left
node.left = node.right
node.right = temp



*/

func invertBinaryTreeBFS(node *TreeNode) *TreeNode {

	if node == nil {
		return nil
	}

	myQueue := internal.InitQueue(10)
	myQueue.Add(node)

	for myQueue.GetSize() != 0 {
		node := myQueue.GetTail().(TreeNode)

		temp := node.Left
		node.Left = node.Right
		node.Right = temp

		if node.Left != nil {
			myQueue.Add(node.Left)
		}

		if node.Right != nil {
			myQueue.Add(node.Right)
		}

	}

	return node

}

/*
use bottom-up appraoch with post order

1

2 3

dfs(1)

 dfs(2), return 0
 dfs(3, return 0)



*/

func diameterOfBinaryTree(node *TreeNode) int {

	res := 0

	var dfsHelper func(*TreeNode) int
	dfsHelper = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		left := dfsHelper(node.Left)
		right := dfsHelper(node.Right)

		res = findMax(res, left+right)

		return 1 + findMax(left, right)

	}
	dfsHelper(node)

	return res

}

/*
Check if balance tree,

(left-right) < 2

base case, if node is nil, return 0
find out the left depthest path and right depthest path



*/

func isBalanced(node *TreeNode) bool {

	res := 0
	var dfsHelper func(*TreeNode) int
	dfsHelper = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		left := dfsHelper(node.Left)
		right := dfsHelper(node.Right)
		res = findMax(res, getAbs(left, right))
		return 1 + findMax(left, right)

	}
	dfsHelper(node)

	if res > 1 {
		return false
	} else {
		return true
	}

}

func isBalanced2(root *TreeNode) bool {

	res := 0
	var dfsHelper func(root *TreeNode) int
	dfsHelper = func(node *TreeNode) int {

		if root == nil {
			return 0
		}

		left := dfsHelper(node.Left) + 1
		right := dfsHelper(node.Right) + 1

		res = findMax(res, getAbs(left, right))

		return findMax(left, right)

	}

	dfsHelper(root)

	if res > 1 {
		return false
	} else {
		return true
	}

}

func getAbs(a int, b int) int {
	x := a - b

	if x < 0 {
		return -x
	} else {
		return x
	}
}

/*
Use post order

	if node1 == nil && node2 == nil{
		return true
	} else if node1 == nill {
		return false
	} else if node2 ==nill{
		return false
	} esle node1.val != node2.val{
		return false
	}
	leftB = isSameTree2(node1.left, node2.left)
	rightB = isSameTree2(node1.right, node2.right)
*/
func isSameTree2PostOrder(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	leftB := isSameTree2PostOrder(p.Left, q.Left)
	rightB := isSameTree2PostOrder(p.Right, q.Right)

	var res bool
	if p.Val == q.Val {
		res = true
	} else {
		res = false
	}

	return res && leftB && rightB

}

func isSameTree2PreOrder(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	left := isSameTree2PreOrder(p.Left, q.Left)
	right := isSameTree2PreOrder(p.Right, q.Right)

	return left && right

}

/*
1
2 3
12 + 13 = 25
1 2
sum = 0
sum = sum*10 + node.val
Use top down appraoch, pre-order
root, left, right


*/

func sumRootLeaf(root *TreeNode) int {

	var dfs func(*TreeNode, int) int

	dfs = func(node *TreeNode, sum int) int {

		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return sum*10 + node.Val
		}
		sum = sum*10 + node.Val
		left := dfs(node.Left, sum)
		right := dfs(node.Right, sum)
		return left + right

	}

	return dfs(root, 0)

}

/*
	 1
	2   3

4 5

	2

4 5

appraoch,
find if node in root same as root of subroot,
if yes, then just check if they are same tree.
*/
func subTreeOfAnotherTree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}

	if isSameTree2PreOrder(root, subRoot) {
		return true
	}

	l := subTreeOfAnotherTree(root.Left, subRoot)
	r := subTreeOfAnotherTree(root.Right, subRoot)

	return l || r

}

/*
     5
  3   8
 1 4  7 9

p = 4, q =3 , ans = 5, it's a BST
p and q are there
top=down pre order appraoch
1. LCA-root =< p or q
2. LCA-root >= p or q
3. p<LCA-root<q


if root == nil {
	return nil
}

if p < root.val && q > root.val {
	//root is the LCA
	return root
} else if p < root.val && q < root.val{
	root = dfs(root.left, p, q)
} else if p > root.val && q > root.val{
	root =eturn dfs(root.right, p, q)
}

return root

*/

func lcaBST(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val > root.Val {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lcaBST(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lcaBST(root.Right, p, q)
	}

	return root

}

/*
pre-oreder, top down


8

1 2
4

if root == p || q == root {
	return root
}

l = dfs(root.left, p, q)
right = dfs(root.left, p, q)
return root


*/

func lcaNonBST(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}
	l := lcaNonBST(root.Left, p, q)
	r := lcaNonBST(root.Right, p, q)

	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	} else {
		return r
	}

}

/*
 	 1
	2 3
  4 5  6 7
level order
[[1]
[2,3]
[4,5,6,7]]

Init a 2-D array to store the final ans
init a queue
put root in queue

for queue is not empty{

	init a list
	for i in range queue size {
		item = queue.poll()
		list.add(item)
		if item.left{
			queue.put(tem.left)
		}

		if item.right{
			queue.put(tem.right)
		}
	}
	asn = append(ans, list)

}



*/

func btLevelOrderTravese(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		var temp []int
		for i := range len(queue) {
			item := queue[0]
			queue = queue[1:]
			temp = append(temp, item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}

			if item.Right != nil {
				queue = append(queue, item.Right)
			}
			i++

		}
		ans = append(ans, temp)
	}

	return ans

}

// BFS
func rightViewTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[0]
			queue = queue[1:]
			if i == size-1 {
				ans = append(ans, item.Val)
			}
			if item.Left != nil {
				queue = append(queue, item.Left)
			}

			if item.Right != nil {
				queue = append(queue, item.Right)
			}

		}
	}

	return ans
}

/*
Use global size and current depth to determine the action
*/
func rightViewTreeDFS(root *TreeNode) []int {
	var ans []int

	var dfsFunc func(root *TreeNode, depth int)

	dfsFunc = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == len(ans) {
			ans = append(ans, root.Val)
		}

		dfsFunc(root.Right, depth+1)
		dfsFunc(root.Left, depth+1)
	}

	dfsFunc(root, 0)

	return ans

}

/*
  2

 1  1
3     5
4
2
11
3 1 5


2 1 3 4 10 2 1
use DFS appraoch , probably pre-order traversal
Traverse the tree to find if current > max, if yes, increase , except the root


dfs(root, max) TreeNode{
	if root == nil {
		return nil
	}

	if root.Val > max{
		incerase
	} else {
		return root
	}
	max = findMax(max, root.Val)

	dfs(root.left, max)
	dfs(root.right, max
    return root

}

  2

 1  1
3     5
*/

func goodNodesTree(root *TreeNode) int {

	count := 0

	max := math.MinInt
	var dfs func(root *TreeNode, max int) int
	dfs = func(root *TreeNode, max int) int {
		if root == nil {
			return 0
		}

		if root.Val >= max {
			count++
		}

		max = findMax(max, root.Val)

		dfs(root.Left, max)
		dfs(root.Right, max)

		return 123123
	}

	dfs(root, max)
	return count

}

/*
in order traverse,
matain a max varibale, if current.val < max then return false,
 2
1 3
if root == nil {
	return true
}

dfs(root.left, root.val)

if root.val < previous {
	return false
}

dfs(root.right)



 2
1 3

 2

1

Cannot really use if previousVal is bigger than current value
*/

func isValidBST2(root *TreeNode) bool {

	var dfs func(root *TreeNode, leftVal int, rightVal int) bool
	dfs = func(root *TreeNode, leftVal int, rightVal int) bool {

		if root == nil {
			return true
		}
		if root.Val > leftVal || root.Val < rightVal {
			return false
		}

		left := dfs(root.Left, root.Val, rightVal)

		right := dfs(root.Right, leftVal, root.Val)

		return left && right

	}

	return dfs(root, math.MaxInt, math.MinInt)

}

func print_list_recursive(head *ListNode, previous int) bool {
	if head == nil {
		return true
	}
	fmt.Printf("number now is %d, number previous is %d", head.Val, previous)

	if head.Val < previous {
		return false
	}
	return print_list_recursive(head.Next, head.Val)
}

func pow(a int, b int) int {

	if b == 0 {
		return 1
	}

	result := pow(a, b/2) * a

	if b%2 == 1 {
		result = result * a
	}
	return result

}

/*
   5
 3   6
1 4

stack,

while stack || node
	while root != nil
		stakc.push(root)
		root = root.left

	node := stack.pop()
	//print
	node = node.right



*/

func kSmallestBST(root *TreeNode, k int) int {

	var dfs func(root *TreeNode, k int, r *int) int
	count := 0
	result := -1
	dfs = func(root *TreeNode, k int, r *int) int {

		if root == nil {
			return 0
		}

		dfs(root.Left, k, r)

		// print root.val

		// fmt.Printf("%v ", root.Val)
		count++
		if count == k {
			*r = root.Val
			return 0
		}

		dfs(root.Right, k, r)

		return 0

	}

	dfs(root, k, &result)
	return result

}

/*

 2
1 3 pre[2,1,3], in-order[1 2 3]
root is pre[0]

root = TreeNode(pre[0])
rootIndexInOrder = inorder.index(pre[0]), rootIndexInOrder indicates how many items in the left and right
root.left = getLeft(pre[1:rootIndexInOrder+1], inorder[:rootIndexInOrder])
root.right = getRight(pre[rootIndexInOrder+1:],inorder[:rootIndexInOrder+1:] )
return root


linkedlist insert recursion


insertList(val, head) {

	if head == nill {
		newNode = ListNode(val)
		return newNode
	}

	head.next = insertList(val, head.next)
}


*/

func buildTree(preOrder []int, inOrder []int) *TreeNode {
	return nil
}

/*
If val is negative, then assign it to 0
2 steps:
1. Get left or right result
2. add together to the current node val
3. do some compare
4. max is a global variable

if root == nil {
	return 0
}

left = dfs(root.left)
right = dfs(root.right)
sum = left + right + root.Val
max = (sum, max)
return max(left + root.val, right + root.val)





*/
// func maxPathSum2(root *TreeNode) int {

// 	var dfs func(root *TreeNode) int
// 	dfs = func(root *TreeNode) int {
// 		if root == nil{
// 			return 0
// 		}

// 	}
// }

/*
BST
    7
  3   8
1  4
 root, p, q

 if root == nil {
	return nil
 }

 if root == p && root != nil {
	return p
 }

 if root == q && root != nil {
	return q
 }

 if p.Val < root.Val || q.Val < root.Val{
	left  = dfs(root.Left, p, q)
 }


 if p.val > root.Val || q.val > root.Val {
	right = dfs(root.right, p, q)
 }

 return root || (left || right)



***************************

Non BST

    1
 3     4
5 6
if root == nil {
	return nil
}

if root == q || root == p
	return root

left = dfs(root.left, p, q)
right = dfs(root.left, p, q)

if left return left



*/

func courseScheduler(numCourses int, prerequisites [][]int) bool {
	preReqMap := make(map[int][]int)
	visitedCrs := make(map[int]int)

	for _, val := range prerequisites {
		preReqMap[val[0]] = append(preReqMap[val[0]], val[1])
	}

	var dfs func(crs int, vistedCrs map[int]int) bool
	dfs = func(crs int, vistedCrs map[int]int) bool {

		if visitedCrs[crs] == 1 {
			return false
		}

		if visitedCrs[crs] == -1 {
			return true
		}

		visitedCrs[crs] = 1

		for _, crs := range preReqMap[crs] {

			if dfs(crs, vistedCrs) == false {
				return false
			}

		}

		visitedCrs[crs] = -1

		return true

	}

	for i := 0; i < numCourses; i++ {
		// visitedCrs[i] = 1
		if !dfs(i, visitedCrs) {
			return false
		}
		// visitedCrs[i] = -1
	}

	return true

}

/*
Fib number
 0 1 1 2 3 5 8
 0 1 2 3 4 5 6

 f(0) = 0
 f(1) = 1


 f(6) = f(5) + f(4)
 	f(5) = f(4) + f(1)
		f(4) = f(3)+f(1)


	f(4) = f(3)+f(1)

global cache
fib(n)

if n in cache:
return cache[n]

if n < 2 {
	return n
}

result = fib(n-1) + fib(n-2)
cache(n) = result

return result

*/

func fib(n int) int {

	var fib func(nn int) int
	globalCache := make(map[int]int)

	fib = func(n int) int {

		if _, ok := globalCache[n]; ok {
			return globalCache[n]
		}
		if n == 0 {
			return 0
		}

		if n == 1 {
			return 1
		}

		result := fib(n-1) + fib(n-2)
		globalCache[n] = result
		return globalCache[n]

	}
	return fib(n)

}

func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(left int, right int, track string)

	dfs = func(left int, right int, track string) {
		if len(track) == 2*n {
			res = append(res, track)
			return
		}

		if left < n {
			track = track + "("
			dfs(left+1, right, track)
			track = track[:len(track)-1]
		}

		if left > right {
			track = track + ")"
			dfs(left, right+1, track)
			track = track[:len(track)-1]
		}

	}

	dfs(0, 0, "")
	return res
}

/*
Input [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]




*/

func subsets(nums []int) [][]int {

	result := [][]int{}
	var dfs func(nums []int, temp []int, i int)
	dfs = func(nums []int, temp []int, j int) {

		subsetCopy := make([]int, len(temp))
		copy(subsetCopy, temp)
		result = append(result, subsetCopy)

		for i := j; i < len(nums); i++ {
			temp = append(temp, nums[i])
			dfs(nums, temp, i+1)
			temp = temp[:len(temp)-1]
		}

	}

	dfs(nums, []int{}, 0)

	return result

}

/*

Input:
nums = [2,5,6,9]
target = 9

Output: [[2,2,5],[9]]

2  9
2
5

base cases:
results[]int
backtracking(sums, target, )

    temp = 0
	for val in sums:{
		temp+=val
	}
	if temp == target {
		append sums to results
		return
	}

	for i in range(nums){
		append nums[i] to sum
		backtrackingsum, target)
		delete nums[i] in sum
	}


}


*/

func combinationSumCanUseSameNumber(candidates []int, target int) [][]int {

	var results = [][]int{}

	var backtracking func(start int, nums []int, target int, sums []int)
	backtracking = func(start int, nums []int, target int, sums []int) {
		// temp := 0

		if target < 0 {
			return
		}

		if target == 0 {
			tempSlice := make([]int, len(sums))
			copy(tempSlice, sums)
			results = append(results, tempSlice)
			return

		}

		for i := start; i < len(nums); i++ {
			sums = append(sums, nums[i])

			backtracking(i, nums, target-nums[i], sums)

			sums = sums[:len(sums)-1]
		}
		// for _, val := range nums {
		// 	sums = append(sums, val)

		// 	backtracking(0, nums, target, sums)

		// 	sums = sums[:len(sums)-1]
		// }

	}
	backtracking(0, candidates, target, []int{})
	return results

}

func combinationSum2UniqueNumber(candidates []int, target int) [][]int {

	var results = [][]int{}

	sort.Ints(candidates)

	// memCache := make(map[]bool, 10)

	var backtracking func(start int, nums []int, target int, sums []int)
	backtracking = func(start int, nums []int, target int, sums []int) {

		if target < 0 {
			return
		}

		if target == 0 {
			tempSlice := make([]int, len(sums))
			copy(tempSlice, sums)
			results = append(results, tempSlice)
			return

		}

		for i := start; i < len(nums); i++ {

			// this is basically in the 2nd big loop
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// if i > 0 && nums[i] == nums[i-1] {
			// 	continue
			// }

			sums = append(sums, nums[i])

			backtracking(i+1, nums, target-nums[i], sums)

			sums = sums[:len(sums)-1]
		}

	}
	backtracking(0, candidates, target, []int{})
	return results

}

/*

Input: coins = [1,5,10], amount = 12

Output: 3

base case, amount == 0 return , amount < 0

memo[amount] = val, val is smallest number of combination plus amount to target

0    1    2    3    4    5    6    7    8    9    10    11    12
          1                   2     1              2      1
*/

func coinChangeRec(coins []int, amount int) int {

	memo := make(map[int]int)

	var backtracking func(coins []int, amount int, collection []int) int
	backtracking = func(coins []int, amount int, collection []int) int {

		if amount < 0 {
			return -1
		}

		if amount == 0 {
			return 0
		}

		if _, ok := memo[amount]; ok {
			return memo[amount]
		}

		minCnt := math.MaxInt
		for i := 0; i < len(coins); i++ {
			collection = append(collection, coins[i])
			res := backtracking(coins, amount-coins[i], collection)
			collection = collection[:len(collection)-1]
			if res == -1 {
				continue
			}
			if res < minCnt {
				minCnt = res + 1
			}
		}
		// This is like checking the root tree after it finished all the childs
		if minCnt == math.MaxInt {
			memo[amount] = -1
		} else {
			memo[amount] = minCnt

		}
		return memo[amount]
	}
	fmt.Print("asd")
	rt := backtracking(coins, amount, []int{})
	return rt

}

func coinChangeCheckDepth(coins []int, amount int) int {
	globalCnt := math.MaxInt
	memo := make([]int, amount+1)

	// Initialize memo: fill with MaxInt except for base case 0
	for i := range memo {
		memo[i] = math.MaxInt
	}
	memo[0] = 0 // base case: 0 coins to make amount 0

	var backtracking func(amount int, depth int)
	backtracking = func(amount int, depth int) {
		if amount < 0 {
			return
		}

		if depth >= memo[amount] {
			// Already found a shorter or equal way to this amount
			return
		}
		memo[amount] = depth

		if amount == 0 {
			if depth < globalCnt {
				globalCnt = depth
			}
			return
		}

		for _, coin := range coins {
			backtracking(amount-coin, depth+1)
		}
	}

	backtracking(amount, 0)

	if globalCnt == math.MaxInt {
		return -1
	}
	return globalCnt
}

func fib3(n int) int {
	cache := make(map[int]int)

	var fib = func(n int) int {

		if _, ok := cache[n]; ok {
			return cache[n]
		}
		if n < 2 {
			return 1
		}

		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}

	return fib(n)

}

func combinationSumAny(candidates []int, target int) bool {

	var backtracking func(candidates []int, target int) bool
	cache := make(map[int]bool)

	backtracking = func(candidates []int, target int) bool {

		if _, ok := cache[target]; ok {
			return true
		}
		if target < 0 {
			return false
		}
		if target == 0 {
			return true
		}

		for i := 0; i < len(candidates); i++ {
			result := backtracking(candidates, target-candidates[i])
			cache[target] = result
			if result {
				return true
			}
		}

		return false

	}
	return backtracking(candidates, target)

}

func combinationSum4(nums []int, target int) int {

	var backtracking func(candidates []int, target int) int
	// global := 0
	cache := make(map[int]int)
	// results := [][]int{}

	backtracking = func(candidates []int, target int) int {

		if _, ok := cache[target]; ok {
			return cache[target]
		}

		if target == 0 {
			return 1
		}

		if target < 0 {
			return 0
		}

		count := 0
		for i := 0; i < len(candidates); i++ {
			count = backtracking(candidates, target-candidates[i]) + count

		}
		cache[target] = count
		return count

	}
	cc := backtracking(nums, target)
	return cc

}

func pow2(x int, n int) int {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	return pow2(x, n-1) * x
}

func comboSumShortrest(nums []int, target int) int {

	var dfs func(nums []int, target int, temp []int) int

	maxNum := math.MaxInt8

	dfs = func(nums []int, target int, temp []int) int {
		if target < 0 {
			return 0
		}

		if target == 0 {
			// get lenth
			if len(temp) < maxNum {
				maxNum = len(temp)
			}
			return 0
		}

		for i := 0; i < len(nums); i++ {
			temp = append(temp, nums[i])
			dfs(nums, target-nums[i], temp)
			temp = temp[:len(temp)-1]
		}
		return 0
	}

	dfs(nums, target, []int{})
	return maxNum

}

/*
climbing stairs

if target < 0 {
	return
}
if target == 0 {
	cnt++
}


for i =1 ;i<=2;i++ {
	dfs(target - i)
}




*/

func climbingStairsBC(n int) int {

	memo := make(map[int]int)
	var dfs func(arget int) int
	dfs = func(target int) int {

		if target < 0 {
			return 0
		}
		if _, ok := memo[target]; ok {
			return memo[target]
		}

		if target == 0 {
			return 1
		}

		// var steps int
		steps := dfs(target-1) + dfs(target-2)
		// for i := 1; i <= 2; i++ {
		// 	steps += dfs(target - i)
		// }

		// steps = dfs(target-1) + dfs(target-2)
		memo[target] = steps

		return memo[target]

	}
	rt := dfs(n)

	return rt
}

/*
A palindrome is a string that reads the same forward and backward.

If there are multiple palindromic substrings that have the same length, return any one of them.

Example 1:

Input: s = "ababd"

Output: "bab"
*/

// func longestPalindrome(s string) int {

// }

/*
You are given an integer array nums where nums[i] represents the amount of money the ith house has. The houses are arranged in a straight line, i.e. the ith house is the neighbor of the (i-1)th and (i+1)th house.

You are planning to rob money from the houses, but you cannot rob two adjacent houses because the security system will automatically alert the police if two adjacent houses were both broken into.

Return the maximum amount of money you can rob without alerting the police.

Example 1:

Input: nums = [1,1,3,3]

Output: 4

1 house
S1 = max(h1)
2 houses
S2 = max(S1,h2)
3 houses
S3 = max(S2, S1+h3
4 houses
S4 = max()



*/

// func houseRober([]int nums) int {

// }

/*
 */
func houseRoberBCandMemo(nums []int) int {

	// maxNum := math.MinInt8
	memo := make(map[int]int)
	var dfs func(index int, nums []int) int
	dfs = func(index int, nums []int) int {
		if index >= len(nums) {
			return 0
		}

		if _, ok := memo[index]; ok {
			return memo[index]
		}

		rob := nums[index] + dfs(index+2, nums)
		skip := dfs(index+1, nums)
		max := findMax(rob, skip)

		memo[index] = max
		return max
	}

	rt := dfs(0, nums)

	return rt

}

func houseRober2(nums []int) int {

	memo := make(map[int]int)

	var dfs func(nums []int, index int) int
	dfs = func(nums []int, index int) int {

		if _, ok := memo[index]; ok {
			return memo[index]
		}

		if index >= len(nums) {
			return 0
		}

		rob := nums[index] + dfs(nums, index+2)
		skip := dfs(nums, index+1)
		memo[index] = findMax(rob, skip)
		return memo[index]

	}
	if len(nums) > 2 {
		nums1 := nums[:len(nums)-1]
		nums2 := nums[1:]
		rt1 := dfs(nums1, 0)
		memo = make(map[int]int)
		rt2 := dfs(nums2, 0)
		return findMax(rt1, rt2)
	} else {
		return dfs(nums, 0)
	}

}

/*

the house is formed by a binary tree, there is only one entry,
alert will be trigger if two adjecent house is robbed .
2 conditions,
If the current house is robbed, I skip the subtrees,
If not, I rob the sub trees.

if rob, rob = currentNode.value + rob(currentNode.Left.Left, currentNode.Right)
If not rob, not_rob=rob(currentNode.Left, currentNode.Right)
return max(rob, not_rob)

*/

func houseRob3(root *TreeNode) int {

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		if root.Left != nil {

		}

		rob := root.Val + dfs(root.Left.Left) + dfs(root.Right)
		skip := dfs(root.Left) + dfs(root.Right)

		return findMax(rob, skip)
	}

	return dfs(root)

}

func findmaxBranchTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := node.Val + findmaxBranchTree(node.Left)
	right := node.Val + findmaxBranchTree(node.Right)

	return findMax(left, right)
}

/*
Check if string has any Palindrome sub string
bcdc

	for i :=0;i<len(s);i++{
		l, r := i, i

		for l >= 0 and r < len(nums){

			if nums[l] == nums[r] {
				temp.append(nums[l])
			}
			l--
			r++
		}

		r = i + 1

			for l >= 0 and r < len(nums){

			if nums[l] == nums[r] {
				temp.append(nums[l])
			}
			l--
			r++
		}

}

bab
*/
func longestPalindrome(s string) string {
	sLen := len(s)
	// result := make([]string, 10)
	resLen := 0
	res := ""

	for i := 0; i < sLen; i++ {
		for j := i; j < sLen; j++ {
			l, r := i, j

			for s[l] == s[r] && l < r {
				l++
				r--
			}

			if l >= r && resLen < j-1+1 {
				resLen = j - i + 1
				res = s[i : j+1]
			}

		}

	}

	return res

}

/*
Max product
Input: nums = [1,2,-3,4]

Output: 4


*/

func maxProduct(nums []int) int {

	minNum := math.MinInt32
	var dfs func(nums []int, start int, sum int)
	dfs = func(nums []int, start int, sum int) {

		if start >= len(nums) {
			return
		}

		if sum > minNum {
			minNum = sum
		}

		for i := start; i < len(nums); i++ {
			if i > start {
				continue
			}
			sum = nums[i] * sum
			dfs(nums, start+1, sum)
			sum = 1
		}

	}
	dfs(nums, 0, 1)
	return minNum
}

/****************
2nd time
neetcode 250
***************/

/*
nums1 =  1 2 3 4 _ _, m
nums2 = 2 6, n

	1 2 3 4 _ 6
	1 2 3 4 4 6
	1 2 3 3 4 5

for item in nums1:

	   if nums1[m-1] < nums1[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
	   } else {
		   nums1[m+n-1] = nums1[m]
		   m--
	   }
*/
func mergeSortedArr(nums1 []int, nums2 []int, m int, n int) []int {

	for n > 0 {
		if m != 0 && nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	return nums1

}

/*

nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

sliding window appraoch
i,j = 0, 1

1011




*/

func containsNearbyDuplicate(nums []int, k int) bool {
	myMap := make(map[int]int)

	for index, val := range nums {

		if _, ok := myMap[val]; ok {
			if index-myMap[val] < k {
				return true
			} else {
				myMap[val] = index
			}

		} else {
			myMap[val] = index
		}

	}

	return false

}

/*
Input: prices = [10,5,1,7,1]

Output: 6
Explanation: Buy prices[1] and sell prices[4], profit = 7 - 1 = 6.
maxProfit = minInt
i = 0, j = 1


for j < len(prices)
	if prices[j] - prices[i] < 0 {
		i++
		continue
	} else {
		max = findmax(max, prices[j] - prices[i])
		j++
	}





*/

/*
Given a string s, find the length of the longest substring without duplicate characters.

A substring is a contiguous sequence of characters within a string.

Example 1:

Input: s = "z x y z d b c"
Input: s = "z x y y a b w"
Input: s = "a a a "

i,j=0, 0

for j < len(nums)

	if map[nums[j]]  {
		delete()

	} else {
		map[nums[j]] = 1
		max = findmax(max, j - i + 1)
		j++
	}

Output: 3
*/
func lengthOfLongestSubstring2(s string) int {
	myMap := map[byte]bool{}
	max := 1

	i, j := 0, 0
	for j < len(s) {
		if _, ok := myMap[s[j]]; ok {
			delete(myMap, s[i])
			i++
		} else {
			myMap[s[j]] = true
			max = findMax(max, j-i+1)
			j++

		}
	}

	return max

}

/*
Example 1:

Input: s = "XYYX", k = 2

Output: 4
Explanation: Either replace the 'X's with 'Y's, or replace the 'Y's with 'X's.

Example 2:

Input: s = "AAABABB", k = 1

Output: 5
XYYX,
}

A A A B| B B B D, k =1
r - l + 1 - maxCnt <= k
*/
func characterReplacement(s string, k int) int {
	myMap := map[byte]int{}
	maxCntKey, l, r := 1, 0, 0
	max := 0

	for r = 0; r < len(s); r++ {
		myMap[s[r]]++
		windlow_len := r - l + 1

		if myMap[s[r]] > maxCntKey {
			maxCntKey = myMap[s[r]]
		}

		if windlow_len-maxCntKey > k {
			myMap[s[l]]--
			l++
		}

		max = findMax(max, r-l+1)

	}
	return max

}

/*

Input: target = 15, nums = [1,2,3,4,5]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

find if sum equal or greater than the target,
	if yes, get minLen, sum = sum - left, left ++

*/

func minSubArrayLen(target int, nums []int) int {
	left, right, minLen, sum := 0, 0, math.MaxInt, 0

	for right = 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {

			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}

	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen

}

/*


 */

func generateParenthesis2(n int) []string {
	res := make([]string, 10)
	// temp := make([]string, 10)
	var dfs func(left int, right int, temp string)
	dfs = func(left int, right int, temp string) {
		if len(temp) == 2*n {
			res = append(res, temp)
		}

		if left < n {
			temp = temp + "("
			dfs(left+1, 0, temp)
			temp = temp[:len(temp)-1]

		}

		if right < n {
			temp = temp + ")"
			dfs(0, right+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, 0, "")
	return res

}
func inorderTraversal2(root *TreeNode) []int {

	res := []int{}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		res = append(res, root.Val)
		traverse(root.Right)
	}

	traverse(root)
	return res

}

func invertBinaryTree3(root *TreeNode) *TreeNode {

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {

		if root == nil {
			return
		}

		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return root

}

/*
BFS appraoch

queue

queue.add(root)

while !queue.empty {
	for lenth in queue{
		node = queue.dequeue()
		swap(node.left. right)
		queue.enqueue(node.left)
		queue.enqueue(node.right)

	}
}

*/

func basicImplementations() {

	userAccount := make(map[string]map[string]interface{})

	if userAccount["123"] == nil {
		userAccount["123"] = make(map[string]interface{})
	}
	userAccount["123"]["name"] = "george"
	userAccount["123"]["transcation"] = 3
	userAccount["123"]["amount"] = 100
	userAccount["123"]["timestamp"] = []int{1, 2, 3}

	if userAccount["456"] == nil {
		userAccount["456"] = make(map[string]interface{})
	}
	userAccount["456"]["name"] = "george"
	userAccount["456"]["transcation"] = 2
	userAccount["456"]["amount"] = 101
	userAccount["456"]["timestamp"] = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("my map value is %v", userAccount)

	// aa := userAccount["456"]["timestamp"].([]int)

	var result []map[string]interface{}
	for userId, info := range userAccount {
		entry := make(map[string]interface{})
		entry["userId"] = userId
		for k, v := range info {
			entry[k] = v
		}
		result = append(result, entry)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i]["transcation"].(int) > result[j]["transcation"].(int) // Sort in ascending order
	})

	// fmt.Printf("done %v\n", mm)

	fmt.Printf("done ")

}
