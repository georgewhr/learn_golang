package main

import (
	"container/heap"
	"fmt"
	"george/internal"
	"math"
	"sort"
	"strconv"
	"strings"
)

type ListNode = internal.ListNode
type List = internal.LinkedList
type PriorityQueue = internal.PriorityQueue
type Stack = internal.Stack
type TreeNode = internal.TreeNode
type TBHS = internal.TBHT
type MinHeap = internal.MinHeap
type Test struct {
	name string
	size int
}

type Human struct {
	name string
	age  *int
}

type File struct {
	child    []*File
	checksum uint64
	name     string
	date     uint64
}

type Node struct {
	Val      int
	Children []*Node
}

func test123(arr []int) {

	arr = append(arr, 2)

}

func main() {

	// t := []int{4, 9}

	// fmt.Print(len(t[3:]))

	arr := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}
	createBinaryTree(arr)
	tt1 := internal.NewIntegerContainerImpl()
	tt1.Add(1)
	wordBreak("applefanap", []string{"fan", "apple", "ap"})

	palindromeSubstringDP("aaa")

	lisOwn([]int{4, 10, 4, 3, 8, 9})

	lisDP([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})

	test := fibTopDown(10)
	fmt.Println(test)
	test = fibBottomUp(10)
	fmt.Println(test)
	uniquePathsBacktracking(10, 8)

	decodeways("226")
	climbStairDFS(3)
	removeOtherPrix([]string{"abc", "ab", "abce", "apple", "app"})

	trieNodeHead := &internal.TrieTable{}
	trieNodeHead.Insert("abc")
	trieNodeHead.Search("abc")

	basicImplementations()
	var head *ListNode
	myList := List{}
	head = myList.InsertRecursion(head, 1)
	head = myList.InsertRecursion(head, 2)
	head = myList.InsertRecursion(head, 3)
	head = myList.InsertRecursion(head, 4)
	// myList.InsertRecursion(myList.Head, 3)
	// myList.Insert(10)
	// myList.Insert(6)
	// myList.Insert(3)
	// myList.Insert(2)

	// myList2 := List{}
	// myList2.Insert(13)
	// myList2.Insert(11)
	// myList2.Insert(8)

	// mergeKsortedListPQ([]*ListNode{myList.Head, myList2.Head})

	s := "george"
	y := s[0]
	fmt.Print(string(y))

	for _, val := range s {
		fmt.Print(string(val))
	}

	basicCalculator("(1+1)*3+1")
	tt := make([]int, 10)
	tt = append(tt, 10)

	fmt.Print(len(tt))

	letterCombinations("34")

	islandPerimeter2([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}})
	// islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}})
	leastInterval([]int{2, 2, 3, 4, 5}, 3)
	minHeap := internal.MinHeapConstructor(10)
	minHeap.InsertVal(10)
	minHeap.InsertVal(11)
	minHeap.InsertVal(5)
	minHeap.InsertVal(7)
	minHeap.InsertVal(2)
	minHeap.InsertVal(3)
	minHeap.InsertVal(1)

	minHeap.RemoveVal()
	minHeap.RemoveVal()
	minHeap.RemoveVal()
	minHeap.RemoveVal()
	minHeap.RemoveVal()

	maxHeap := internal.MaxHeapConstructor(10)
	maxHeap.InsertVal(10)
	maxHeap.InsertVal(11)
	maxHeap.InsertVal(5)
	maxHeap.InsertVal(7)
	maxHeap.InsertVal(2)
	maxHeap.InsertVal(3)
	maxHeap.InsertVal(1)

	maxHeap.RemoveVal()
	maxHeap.RemoveVal()
	maxHeap.RemoveVal()
	maxHeap.RemoveVal()
	maxHeap.RemoveVal()

	rotateArr([]int{1, 2, 3, 4}, 2)
	testtest := make([]int, 10)
	testtest = append(testtest, 2)

	test1 := make([]*ListNode, 12)
	// v := test1[2]
	if test1[1] == nil {
		fmt.Print(test1)
	}

	ag := []int{1, 2}
	test123(ag)

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
	employee.Register("george", "11")
	employee.Register("george", "12")

	employee.Promote("george", "S", "200", "15")

	employee.Register("george", "13")
	employee.Register("george", "14")
	employee.Register("george", "16")
	employee.Register("george", "17")
	employee.Register("george", "18")

	employee.CalculateSalary("george", 10, 20)
	employee.TopN("10", "J")

	cloudStorage := internal.Init()
	inMemDb := internal.InitInDB()

	inMemDb.SetAtWithTTL("A", "T", "G", "3", "5")
	inMemDb.SetAtWithTTL("A", "B", "C", "4", "14")
	inMemDb.SetAtWithTTL("A", "D", "E", "5", "15")
	inMemDb.Scan("A")
	inMemDb.Backup("4")
	inMemDb.Backup("12")
	inMemDb.Restore("100", "12")

	// cloudStorage.Add("/root/file2.mp3", "56")
	// cloudStorage.Add("/root/file1.mp3", "20")
	cloudStorage.Add("/root/file1.mp3", "10")
	cloudStorage.Add("/root/file1.mp3", "40")
	cloudStorage.Add("/root/file1.mp3", "50")
	cloudStorage.Add("/root/file1.mp3", "60")

	cloudStorage.Add("/root/file2.mp3", "20")
	cloudStorage.Add("/root/file2.mp3", "220")
	cloudStorage.Add("/root/file2.mp3", "400")
	cloudStorage.Add("/root/file2.mp3", "600")
	// cloudStorage.DeleteVersion("/root/file1.mp3", "0")
	cloudStorage.TopN("/root/file", "2")
	// cloudStorage.Copy("/dir100/file100.mp3", "/dir100/copy/file100.mp3")
	// cloudStorage.Add("/dir/file2.txt", "20")
	// cloudStorage.Add("/dir3/file3.txt", "50")

	// cloudStorage.FindFile("/root", ".mp3")
	// cloudStorage.AddUsers("george", "100")
	// cloudStorage.AddFileBy("george", "/dir/file2.txt", "40")
	// cloudStorage.Copy("/dir/file2.txt", "/dircopy/file2.tx")
	// cloudStorage.GetFileSzie("/dir/file2.txt")

	// // cloudStorage.AddUsers("vera", "200")
	// cloudStorage.AddFileBy("george", "/dir101/file100.mp3", "40")
	// cloudStorage.AddFileBy("george", "/dir101/file1001.mp3", "50")
	// cloudStorage.Copy("/dir101/file1001.mp3", "/dir100/copy1/file100.mp3")
	// cloudStorage.UpdateCapacity("george", "50")

	// bank := internal.InitBankingSystem()
	// bank.CreateAccount(100, "george")
	// bank.CreateAccount(120, "vera")
	// bank.Deposit(130, "george", 100)
	// bank.Deposit(130, "vera", 200)
	// bank.Pay("200", "george", 20)
	// bank.Pay("201", "vera", 15)
	// bank.Pay("202", "george", 40)
	// bank.Transfer(300, "george", "vera", 5)
	// bank.Withdraw(400, "george", 10)
	// bank.GetPaymentStatus(86500000, "george", "george_0")

	// cloudStorage.AddFileBy("george", "/dir102/file1001.mp3", "50")

	// test := []string{"cir", "car"}
	// test := []string{"fla", "flw"}
	// test1 := "abc"
	// fmt.Printf("number is %s", test1[0:1])
	// output := longestCommonPrefix(test)
	// fmt.Printf("number is %f\n", output)

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

	constructBTPreInorder([]int{1, 2, 3}, []int{2, 1, 3})
	invertBinaryTreeBFS2(rootNode)
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
	maxProductDPTopDown([]int{2, 3, -2, 4})
	maxProduct([]int{1})
	wordLadder([]string{"hot", "dot", "dog", "lot", "log", "cog"}, "hit", "cog")

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

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
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

func fibTopDown(n int) int {

	dp := make(map[int]int)
	var dfs func(n int) int
	dfs = func(n int) int {

		if _, ok := dp[n]; ok {
			return dp[n]
		}
		if n == 0 || n == 1 {
			return 1
		}

		res := dfs(n-1) + dfs(n-2)
		dp[n] = res
		return dp[n]
	}

	return dfs(n)

}

func fibBottomUp(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]

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

func coinChangeRec2NoMemo(coins []int, amount int) int {

	memo := make(map[int]int)

	var backtracking func(coins []int, amount int) int
	backtracking = func(coins []int, amount int) int {

		if amount < 0 {
			return -1
		}

		if amount == 0 {
			return 0
		}

		if _, ok := memo[amount]; ok {
			return memo[amount]
		}
		minCnt := math.MaxInt16

		for i := 0; i < len(coins); i++ {
			rt := backtracking(coins, amount-coins[i])
			if rt == -1 {
				continue
			}

			if rt < minCnt {
				minCnt = rt + 1
			}

		}

		if minCnt == math.MaxInt16 {
			memo[amount] = -1
		} else {
			memo[amount] = minCnt
		}

		return memo[amount]

	}
	memo[amount] = backtracking(coins, amount)
	return memo[amount]

}

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

/*[2, 3 5], 15

2, 15, 13, 11, 9, 7, 5, 3, 1, -1
3      12  9,

 1 2 3  5

dp = make([]int, len(target))
dp[0] = 0
dp[1] = 15
dp[2] = 1
dp[3] = 2
dp[4] = 2
dp[5] = 1
dp[6] = 2
dp[7] = 2
dp[8] = 2
dp[9] = 3
dp[10] = 2
dp[11] = 15
dp[12] = 3
dp[13] = 3
dp[14] = 4

10 - 2 = 8
10 - 3 =7
10 -5 = 5

4-2=2
4-3=1

for i =0 -> n{
	for coin in coins {
		if i - coin < 0{
			continue
		}
		dp[i] = min(dp[i-coin]+1, dp[i])

	}
}
2, 3 5

0 0
1 15
2

*/

func coinChangeBottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i, _ := range dp {
		dp[i] = amount + 1 // can be any big number greater than amount
	}

	dp[0] = 0

	for i := 0; i < len(dp); i++ {

		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			dp[i] = findMin(dp[i], dp[i-coin]+1)
		}

	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}

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
Example 1:

Input: s = "ababd"

Output: "bab"

the idea is to check if dp[i-1][j+1] palindrome, if it is,
then just check if dp[i] == dp[j] then plus 2

for i -> n-1...0 {

		for j = i, j<n;j++{
			if nums[i] == nums[j] {
			dp[i][j] = 2 + dp[i+1][j-1]
		} else {

			dp[i][j] =max( dp[i+1][j], dp[i][j-1]
		}
		}
	}

	for
*/

/*
if dp[i+1][j-1]
dp[i][j] = dp[i+1][j-1]  + 2, if s[i] == s[j]
else
dp[i][j] = max(d[i][j-1], )
*/
func longestPalindromeDPLenth(s string) int {
	n := len(s)
	dp := make([][]int, n)

	for index, _ := range dp {
		dp[index] = make([]int, n)
		dp[index][index] = 1
	}

	for i := n - 1; i >= 0; i-- {

		for j := i; j < n; j++ {

			if (j-i <= 2 || dp[i+1][j-1] == 1) && s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = findMax(dp[i][j-1], dp[i+1][j])
			}

		}

	}

	return 0

}

func longestPalindromeDPSub(s string) string {
	n := len(s)
	dp := make([][]int, n)
	resIds, resLen := 0, 0

	for index, _ := range dp {
		dp[index] = make([]int, n)
		dp[index][index] = 1
	}

	for i := n - 1; i >= 0; i-- {

		for j := i; j < n; j++ {

			if (j-i <= 2 || dp[i+1][j-1] == 1) && s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				if resLen <= j-i+1 {
					resLen = j - i + 1
					resIds = i
				}
			}

		}

	}

	return s[resIds : resLen+resIds]

}

/*

if s[i] == s[j] {
	dp[i][j] =dp[i+1][j-1] + 3
} else {
	dp[i][j] =dp[i+1][j-1] + 2
}


2 pointers
Example 1:

Input: s = "abc"

Output: 3
Explanation: "a", "b", "c".

Example 2:

Input: s = "aaa"

Output: 6


for i 0->n-1
   p1, p2 = i
   for p1 >=0 && p2 < n {
	   if s[p1] == s[p2] {
		 res ++

	   } else {
		 break
	   }
	   p1 --
	   p2 ++
   }

   i++

aaa

a, 1
a  1
1
aaa 1
aa 1
aa 1
*/

func palindromeSubstring2Pointer(s string) int {
	res := 0
	n := len(s)

	p1, p2, p3 := 0, 0, 0
	for p1 < n {
		p2, p3 = p1, p1

		for p2 >= 0 && p3 < n {
			if s[p2] == s[p3] {
				res++

			} else {
				break
			}
			p2--
			p3++
		}

		p2, p3 = p1, p1+1

		for p2 >= 0 && p3 < n {
			if s[p2] == s[p3] {
				res++

			} else {
				break
			}
			p2--
			p3++
		}
		p1++

	}

	return res

}

/*
abcb
aaa
dp = make(n, 1)

if (s[i+1:j-1] || j -i <=2 ) && s[i] == s[j]
s[i:j] = s[i+1:j-1] ? && s[i] == s[j]? , total++
final result is s[0:n-1]
*/
func palindromeSubstringDP(s string) int {
	n := len(s)
	dp := make([][]int, n)
	res := 0

	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := n; i >= 0; i-- {

		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1] == 1) {
				res++
				dp[i][j] = 1

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

/*
Input: intervals = [[1,3],[2,5],[6,7]]
Input: intervals = [[1,3],[1,2],[6,7]]

Output: [[1,5],[6,7]]

1. sort ascending order based on 1st element of array.
2. i =0, j = 1

	for j < len(intervals) {
		if intervals[j][0] <= intervals[i][1] {
			ret.add(merge(intervals[j][0], intervals[i][1]))
			j++

		} else {
			ret.add(merge(intervals[j], intervals[i]))
			i = j
			j++
		}
	}
*/
func MergeInterval(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] > intervals[j][0]
	})

	ret := make([][]int, 0)
	lastend := intervals[0][1]

	i, j := 0, 1

	for j < len(intervals) {
		if intervals[j][0] <= intervals[j][0] {
			lastend = findMax(intervals[j][1], lastend)
		} else {
			ret = append(ret, []int{intervals[i][0], lastend})
			i = j
		}
		j++

	}

	ret = append(ret, intervals[i])
	return ret

}

func MaxDepthTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepthTree(root.Left) + 1
	right := MaxDepthTree(root.Left) + 1

	max := findMax(left, right)
	return max

}

func findDiameter(root *TreeNode) int {

	res := 0

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		res = findMax(res, left+right)

		return findMax(left, right) + 1

	}

	dfs(root)

	return res

}

func checkIfBalacnedTree(root *TreeNode) bool {
	var traverse func(root *TreeNode) int
	res := true
	traverse = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		left := traverse(root.Left) + 1
		right := traverse(root.Right) + 1

		temp := absolute(left, right)

		if temp > 1 {
			res = false
		}
		return findMax(left, right)

	}

	traverse(root)
	return res

}

func absolute(a int, b int) int {
	res := a - b
	if res < 0 {
		return -res
	} else {
		return res
	}
}

/*
edge case

	if p == nil && q == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	if p.val != q.val {
		return false
	}

left = sameTree(p.left, q.left)
right = sameTree(p.right, q.right)

return left && right
*/
func sameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)

}

func subTreeOfAnotherTree2(root *TreeNode, subroot *TreeNode) bool {

	if root == nil && subroot == nil {
		return true
	}

	if root == nil || subroot == nil {
		return false
	}

	temp := false
	if root.Val == subroot.Val {
		temp = subTreeOfAnotherTree2(root.Left, subroot.Left) && subTreeOfAnotherTree2(root.Right, subroot.Right)
	}
	left := subTreeOfAnotherTree2(root.Left, subroot)
	right := subTreeOfAnotherTree2(root.Right, subroot)

	return left || right || temp

}

/*
Post order
edge case:
if root == nil {
	return nil
}

if root == p || root == p{
	return root
}

left:= dfs(root.left)
right:= dfs(root.left

if left !=nil && right != nil{
	return root
}

if left !=nil{
	return left
}

if right !=nil{
	return left
}

return


*/

func lowestCommonAncestor2(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return left
	}
	return nil
}

/*
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)

	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root

}

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		/*
			1. if the node has 2 childen
			2. if the node has only left child
				return left child
			3. if the node has only rihgt child
				return right child
			4. if the node has no child
				do nothing, return nil

		*/
		if root.Left == nil {
			return root.Right

		} else if root.Right == nil {
			return root.Left

		} else if root.Left == nil && root.Right == nil {
			return nil
		} else {
			//find the minimum number from right subtree and replace with it
			root.Val = findMinNumber(root.Right)
			root.Right = deleteNode(root.Right, root.Val)

		}
	}
	return root

}

func findMinNumber(node *TreeNode) int {

	minNum := node.Val

	for node != nil {
		if node.Val < minNum {
			minNum = node.Val
		}
		node = node.Left
	}

	return minNum
}

func getConcatenatio2n(nums []int) []int {

	n := len(nums)
	ans := make([]int, 2*n)

	for i, val := range nums {
		ans[i] = val
		ans[n+i] = val
	}

	return ans

}

func validAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := make(map[rune]int), make(map[rune]int)

	for i, val := range s {
		sMap[val]++
		tMap[rune(t[i])]++
	}

	for k, v := range sMap {
		if v != tMap[k] {
			return false
		}
	}

	return true

}

func twoSum2Asc(numbers []int, target int) []int {

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	left, right := 0, len(numbers)-1

	for left != right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--

		} else {
			return []int{numbers[left], numbers[right]}
		}

	}
	return nil

}

func twoSum2(nums []int, target int) []int {

	myMap := make(map[int]int)

	for i, val := range nums {
		temp2 := target - val
		if _, ok := myMap[temp2]; ok {
			return []int{i, myMap[temp2]}
		}
		myMap[val] = i

	}

	return nil

}

/*

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

*/

func longestCommonPrefix2(strs []string) string {

	res := ""
	if len(strs) == 0 {
		return res
	}

	shortLen := len(strs[0])

	for _, val := range strs {
		if len(val) < shortLen {
			shortLen = len(val)
		}
	}

	for i := 0; i < shortLen; i++ {

		table := make(map[byte]int)
		for _, val := range strs {
			table[val[i]]++

			if table[val[i]] == len(strs) {
				res += string(val[i])
				delete(table, val[i])
			}

			if len(table) > 1 {
				return res
			}
		}

	}

	return res

}

func groupAnagram(strs []string) [][]string {
	res := [][]string{}

	table := make(map[[26]int][]string)

	for _, str := range strs {
		var key [26]int
		for _, char := range str {
			key[char-'a']++
		}
		if _, ok := table[key]; !ok {
			table[key] = make([]string, 0)
		}
		table[key] = append(table[key], str)
	}

	for _, val := range table {

		sameGroup := make([]string, 0)
		for _, item := range val {
			sameGroup = append(sameGroup, item)
		}
		res = append(res, sameGroup)
	}

	return res
}

/*

Input: nums = [1,3,2,3], val = 3
Output: 2, nums = [2,2,_,_]
*/

func removeElement2(nums []int, val int) int {

	read, write := 0, 0
	for read < len(nums) {

		if nums[read] != val {
			nums[write] = nums[read]
			write++
		}
		read++

	}

	return write

}

/*
Input: nums = [1,0,1,2]

Output: [0,1,1,2]
*/
func sortColor2(nums []int) {
	temp := [3]int{0, 1, 2}

	for _, val := range nums {
		temp[val]++
	}

	index := 0
	for i := 0; i < 3; i++ {
		for temp[i] > 0 {
			nums[index] = i
			temp[i]--
			index++
		}
	}

}

/*
 */
func topKFrequentCountingSort(nums []int, k int) []int {

	maxVal := nums[0]

	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		}
	}

	countArr := make([]int, maxVal+1)

	for _, val := range nums {
		countArr[val]++
	}

	index := 0
	for num, freq := range countArr {
		temp := freq
		for temp > 0 {
			nums[index] = num
			temp--
			index++
		}
	}

	return nums

}

/*
Init a Minheap with the comparator
enqueue
*/
// func topKFrequentMinHeap(nums []int, k int) []int {

// 	count := make(map[int]int)

// 	for _, val := range nums {
// 		count[val]++
// 	}

// 	heap := priorityqueue.NewWith(func(a, b interface{}) int {
// 		freqA := a.([2]int)[0]
// 		freqB := b.([2]int)[0]
// 		return utils.IntComparator(freqA, freqB)
// 	})

// }

/*
Input: nums = [2,20,4,10,3,4,5]

Output: 4
*/
func longestConsecutive3(nums []int) int {
	table := make(map[int]int)

	for _, val := range nums {
		table[val]++
	}

	max := 1
	for _, val := range nums {
		if _, ok := table[val-1]; !ok {

			sequence := 1
			for {
				if _, ok := table[val+sequence]; ok {
					sequence++
					continue
				}
				max = findMax(max, sequence)
				break
			}

		}
	}

	return max

}

func isPalindrome2(s string) bool {
	start, end := 0, len(s)-1

	for start != end {

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

/*



 */

func longestPalindrome2(s string) int {
	res := 1

	for i := 0; i < len(s); i++ {

		l, r := i, i

		for l >= 0 && r < len(s) && s[l] == s[r] {
			res = findMax(res, l-r+1)
			l--
			r++
		}

	}

	for i := 0; i < len(s); i++ {

		l, r := i, i+1

		for l >= 0 && r < len(s) && s[l] == s[r] {
			res = findMax(res, l-r+1)
			l--
			r++
		}

	}

	return res

}

func mergeAlternately(word1 string, word2 string) string {

	w1Len := len(word1)
	w2Len := len(word1)

	i := 0
	j := 0
	var sb strings.Builder

	for i < w1Len && j < w2Len {
		sb.WriteByte(word1[i])
		sb.WriteByte(word1[j])
	}

	return ""

}

/*
Input: nums1 = [10,20,20,40,0,0], m = 4, nums2 = [1,2], n = 2

10, 20, 30, 0, 0
10,30
Output: [1,2,10,20,20,40]

i = m -1
j = n -1
nums1EndIndex = len(nums1) - 1

	for i > =0 && j >=0 {
		if nums1[i] > nums2[n]{
			//copy nums1[i] to nums[nums1EndIndex]
			nums1EndIndex --
			i--
		} else {
			//copy nums2[j] to last index of nums1
			nums1EndIndex --
			j --
		}

}
*/
func mergedSortedArr(nums1 []int, m int, nums2 []int, n int) {

}

/*
Input: nums = [1,1,2,3,4]
Output: [1,2,3,4]

i =0, j = 0

	for j < len{
		nums[i] = nums[j]
		for j < len && nums[i] = nums[j] {
			j++
		}
		i++
	}
*/
func removeDuplicatesSorted(nums []int) {

}

/*
Input: nums = [-1,0,1,2,-1,-4]

	-4, -1, -1 , 0, 1, 2

	x + y + z = 0

Output: [[-1,-1,2],[-1,0,1]]

this problem is target is zero, so if the index number is greater than 0, then it's impossible to have
sum equal to zero
*/
func threeSum2Sort(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := make([][]int, 0)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		//Check if number is greater than 0
		if nums[i] > 0 {
			break
		}

		//skip same number
		if i > 0 && nums[i]-1 == nums[i] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left != right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--

			} else {
				res = append(res, []int{i, left, right})
			}
		}

	}
	return res
}

/*
Rotate arr

Input: nums = [1,2,3,4], k = 2, 3 4 1 2
Input: nums = [1,2,3,4], k = 3, 2 3 4 1

swap, index, (index + 2) % size
*/
func rotateArr(nums []int, k int) {
	newArr := make([]int, len(nums))

	for i, _ := range nums {
		newArr[(i+k)%len(nums)] = nums[i]
	}
	nums = newArr
	copy(nums, newArr)

	fmt.Print(newArr)

}

/*
Input: prices = [10,1,5,6,7,1]
1. have a variable to store the max value
2. use 2 pointers to find the max value




Output: 6
*/

func maxProfit2(prices []int) int {
	max := math.MinInt16

	i, j := 0, 1

	for j < len(prices) {

		// has profit
		if prices[j]-prices[i] > 0 {
			max = findMax(max, prices[j]-prices[i])
		} else {
			i = j
		}
		j++

	}

	return max
}

/*

fixed window problem
map[rune]int

map s1 to map

Input: s1 = "abc", s2 = "lecabee"

Output: true
r = 0
l = 0, r = len(s1) -1
for r<len {

	copy map to a new map

	temp = 0
	for i = l, i < len {
		if nums[i] in copymap {
			copymap[nums[i]--
		} else {
			break
		}

		if copymap[nums[i]] == 0 {
			delete
		}
	}

	if copymap len is zero, return true
	r ++


}


*/

func checkInclusion2(s1 string, s2 string) bool {

	l, r := 0, len(s1)-1

	table := make(map[rune]int)
	for _, val := range s2 {
		table[val]++
	}

	for r < len(s2) {

		newTable := make(map[rune]int, len(table))

		for k, val := range table {
			newTable[k] = val
		}
		for i := l; i <= r; i++ {

			if _, ok := newTable[rune(s2[i])]; ok {
				newTable[rune(s2[i])]--

				if newTable[rune(s2[i])] == 0 {
					delete(newTable, rune(s2[i]))
				}
			} else {
				break
			}

		}

		if len(newTable) == 0 {
			return true
		}

	}

	return false

}

func lengthOfLongestSubstring3(s string) int {
	table := make(map[rune]int)

	i, j, max := 0, 0, 1

	for j < len(s) {

		if _, ok := table[rune(s[j])]; ok {
			i++

		} else {
			table[rune(s[j])]++
			max = findMax(j-i, max)
			j++

		}
	}

	return 0

}

/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

thoughts:
1. beginword not ncessaryily in the word list, but endWord will.
2. need function to check if word minum letter is in the word list

# BSF

res = 1
vistedMap
q.add(begineWord)

	for len(q) != 0 {
		qlen = len(q)


		for i:=0i<qlen;i++{
			parent = q.pop
			if parent == endword {
				return res
			}
			for item in wordList{
				if checkIfRemoveLetterWordInList(parent, item) && !vistedMap(item) {
					q.add(item)
					vistedMap[item] = true
				}

			}
		}
	}
*/
func wordLadder(wordList []string, beginWord string, endWord string) int {
	// vistedWord := make(map[string]int)
	worddListMap := make(map[string]int)
	res := 0

	for _, val := range wordList {
		worddListMap[val]++
	}

	q := []string{}

	q = append(q, beginWord)

	for len(q) != 0 {
		qLen := len(q)
		res++
		for i := 0; i < qLen; i++ {
			parent := q[0]
			q = q[1:]

			if parent == endWord {
				return res

			}

			for char := 'a'; char <= 'z'; char++ {
				for i := 0; i < len(parent); i++ {
					if char == rune(parent[i]) {
						continue
					}
					newStr := parent[:i] + string(char) + parent[i+1:]

					if _, ok := worddListMap[newStr]; ok {
						q = append(q, newStr)
						delete(worddListMap, newStr)
					}

				}
			}

		}
	}
	return 0

}

/*
Input: points = [[0,2],[2,0],[2,2]], k = 2

Output: [[0,2],[2,0]]
*/
func kClosest(points [][]int, k int) {
	temp := []float64{}

	for _, val := range points {
		a := float64(val[0])
		b := float64(val[1])
		temp = append(temp, math.Sqrt(a*a+b*b))
	}

	/*initt a max heap
	for val in temp {
		heap.add(val)
		for len(heap) > k {
			heap.Pop()
		}
	}


	}

	*/

}

/*
Input: tasks = ["A","A","A","B","C"], n = 3

Output: 9
A:3
B:1
C:1
maxHeap
item = maxHeap.Pop()
map[item] --
if map[item] == 0, delete item in map

for len(map) != 0 {

	k = 0
	for entry in map {
		k++
		heap.add(entry)
		if k == n + 1 {
			break
		}

	}
	for len(q) < n + 1 && maxHeap.len() > 0
		item = maxHeap.Pop()
		map[item] --
		if map[item] == 0, delete item in map
		q.add(item)
	for len(q) < n + 1 {
		q.add(Idel)
	}

}
*/
func leastInterval(tasks []int, n int) int {
	taskTable := make(map[int]int)
	// maxHeap := internal.MaxHeapConstructor(10)
	taskSchduleMaxHeap := &internal.TaskScheduleMaxHeap{}
	heap.Init(taskSchduleMaxHeap)
	q := make([]internal.TaskSchedule, 0)
	res := ""

	for _, val := range tasks {
		taskTable[val]++

	}

	for len(taskTable) > 0 {
		j := 0

		for k, v := range taskTable {
			// item :=
			heap.Push(taskSchduleMaxHeap, internal.TaskSchedule{ID: k, Priority: v})
			if j == n+1 {
				break
			}
		}

		for len(q) < n+1 && taskSchduleMaxHeap.Len() > 0 {
			item := heap.Pop(taskSchduleMaxHeap).(internal.TaskSchedule)
			taskTable[item.ID]--
			if taskTable[item.ID] == 0 {
				delete(taskTable, item.ID)
			}
			q = append(q, item)
		}

		for len(q) < n+1 {
			q = append(q, internal.TaskSchedule{ID: 999, Priority: 999})
		}

		for len(q) > 0 {
			res = res + strconv.Itoa(q[0].ID) + "  "
			q = q[1:]

		}

	}

	return 0

}

/*
Check if the next cell is out of boundary or in the water
if yes, then increase perimeter by 1
m = len(grid)
n = len(grid[0])

lands = {[0,1], [1,0],[1,1],[1,2],[2,1],[3,0],[3,1]}



dfs = func(spot, lands) {

	if spot is land {
		return
	}

	if spot is out of boundary or spot is water{
		perimeter ++
	}

	if visted[spot] == true {
		return
	}

	visted[spot] = true


	dfs all 4 directions


}





for land in lands {
	dfs(land + 1, grid)

}

Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
Output: 16
Explanation: The perimeter is the 16 yellow stripes in the image above.


*/

func islandPerimeter(grid [][]int) int {

	var dfs func(r int, c int)
	p := 0
	m := len(grid)
	n := len(grid[0])
	visted := make(map[[2]int]bool)

	dfs = func(r int, c int) {
		if r >= m || c >= n || r < 0 || c < 0 || grid[r][c] == 0 {
			p++
			return
		}
		spot := [2]int{r, c}
		if visted[spot] {
			return
		}
		visted[spot] = true

		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}

		}
	}

	return p

}

/*
The idea is to serach all the land srounding land, and marked the serached to 0

dfs(x, y) {

	if x, y out of boundary or is water {
		return
	}

	table[x,y] = 0


	dfs(x+1, y)
	dfs(x-1, y)
	dfs(x, y+1)
	dfs(x, y-1)

}

	for land in islands {
		dfs([x, y])
		island++
	}
*/
func numsIslands(grid [][]byte) int {

	numIslands := 0
	m := len(grid)
	n := len(grid[0])
	var dfs func(r int, c int)

	dfs = func(r int, c int) {

		if r >= m || c >= n || r < 0 || c < 0 || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numIslands++
				dfs(i, j)
			}

		}
	}

	return numIslands
}

func letterCombinations(digits string) []string {

	// str := "ghi"
	res := []string{}

	digitToChar := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(start int, temp string)
	dfs = func(start int, temp string) {

		if len(temp) == len(digits) {
			res = append(res, temp)
			return
		}

		for _, val := range digitToChar[digits[start]] {
			temp += string(val)
			dfs(start+1, temp)
			temp = temp[:len(temp)-1]
		}

	}

	dfs(0, "")

	return res

}

/*
123 + 20
1. a function to convert string to int

res = 0
sign = 1 //1 is plus. -1 is negative
for each char in string {


	if char is number{
		res = char + res*10
	}


	else if char is + - {

		if sign == 1 {
			stack.push(res)
		} else if sign == -1 {
			stack.push(-res)
		}

		res = 0

		if char is - {
			sign = -1
		} else if char is + {
			sign = 1
		}

	}
}


11 + 1
2 *3
*/

func basicCalculator(s string) int {

	var sign byte
	sign = '+'
	res := 0
	stack := make([]int, 0)
	finalRes := 0
	subStart, subEnd := 0, 0

	for index := 0; index < len(s); {
		val := s[index]
		if isDigits(val) {
			res = int(val-'0') + res*10
		}

		if val == '-' || val == '+' || val == '*' || val == '/' || index == len(s)-1 {
			if sign == '+' {
				stack = append(stack, res)
			} else if sign == '-' {
				stack = append(stack, -res)
			} else if sign == '*' {
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp := pre * res
				stack = append(stack, temp)

			} else if sign == '/' {
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp := pre / res
				stack = append(stack, temp)
			}
			sign = val

			res = 0
		}

		if val == '(' {
			subStart = index

			for l := index; l < len(s); l++ {
				if s[l] == ')' {
					subEnd = l
					subSum := basicCalculator(s[subStart+1 : subEnd])
					sign = s[subEnd+1]
					stack = append(stack, subSum)
					index = l + 1
					break
				}
			}

		}
		index++
	}

	for _, val := range stack {
		finalRes += val
	}

	return finalRes

}

func isDigits(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	} else {
		return false
	}

}

/*
Declare a priority q, say , pq

dummyNode
temp = DumyNode

	for headNode := range lists {
		pq.enque(headNode)
	}

	for !pq.empty() {
		item = pq.dequeue()
		temp.Next = item
		temp = temp.next

		pe.enqueue(item.next)

}
*/
func mergeKsortedListPQ(lists []*ListNode) {

	// minHeap := internal.MinHeapConstructor(10)
	minHeap := &internal.MinHeapListNode{}
	heap.Init(minHeap)

	for _, val := range lists {
		heap.Push(minHeap, val)
	}

	for minHeap.Len() != 0 {
		item := heap.Pop(minHeap).(*ListNode)
		fmt.Print("\nnumer is \n", item.Val)
		if item.Next != nil {
			heap.Push(minHeap, item.Next)
		}

	}

}

/*
Use DFS to find the loop
3->[1, 2]
1->3
visted[]

	dfs() {
		if visted[person] {
			return false
		}

}
Input: n = 3, trust = [[1,3],[2,3]]
Output: 3
*/
func findJudge(n int, trust [][]int) int {

	visted := make(map[int]int)
	table := make(map[int][]int)

	for _, val := range trust {
		table[val[1]] = append(table[val[1]], val[0])
	}

	var dfs func(person int, visted map[int]int) bool
	dfs = func(person int, visted map[int]int) bool {

		if visted[person] == 1 {
			return false
		}

		visted[person] = 1

		for _, val := range table[person] {
			return dfs(val, visted)
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if !dfs(i, visted) {
			return 1
		}
	}

	return -1

}

/*
str, st, strg, strgeorge, apple, apple1, apple2
strgeorge, apple1, apple2

str st

st str
for i < len

	for j = i, < len {
		if list[i] and list[j] is prefix for each other
	}

checkIfPrefix
check short lenth,

	for i =0 < len(shortStr) {
		if str1[i] != str2[i] {
			return false
		}
	}

return true

Use trie or nested hashmap
str, st, strg, strgeorge

	for index, str in list {
		indexTable[str] = index
		checkIfPrefix(str, indexTable)
	}

	checkIfPrefix(str,indexTable ) {
		for index, char in str {
			if trie.head.children[char] == nil {
				tire.insert(char)
				tire = trie.next
				if indexTable[arr] {
					remove indexTable(arr)
				}

			} else {
				trie = tire.head.children[char]
			}
		}

		if trie.next != nil {
			remove str
			remove str from indextable
		}
	}

str, st, strg, strgeorge, apple, apple1, apple2
*/
func removeOtherPrix(str []string) {

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {

			if len(str[i]) == 0 || len(str[j]) == 0 {
				continue
			}
			res := checkPrefixEach(str[i], str[j])
			if res == 2 {
				continue
			} else if res == 0 {
				fmt.Println(str[i])
				str[i] = ""
			} else {
				fmt.Println(str[j])
				str[j] = ""
			}

		}
	}

}

func checkPrefixEach(str1 string, str2 string) int {
	if len(str1) < len(str2) {

		for i := 0; i < len(str1); i++ {
			if str1[i] != str2[i] {
				return 2
			}
		}

		return 0

	} else {
		for i := 0; i < len(str2); i++ {
			if str1[i] != str2[i] {
				return 2
			}
		}
		return 1
	}

}

/*
q.add(root)
minimumPath = 1
for q.len() != emptry {
	ql = q.len()
	for i > ql {
		node = q.deque()
		if node.left ==nil && node.right == nill {
			minimumPath ++
			return minimumPath
		}

		if node.left != nil {
			q.enqueue{node.left}
		}

		if node.right != nil {
			q.enqueue{node.right}
		}

	}
}


*/

func findMinimumPathOfTreeBFS(root *TreeNode) int {

	if root == nil {
		return 0
	}
	q := []*TreeNode{}
	q = append(q, root)
	minimumPath := 0

	for len(q) != 0 {
		ql := len(q)
		minimumPath++

		for i := 0; i < ql; i++ {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				return minimumPath
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

		}
	}

	return minimumPath

}

func findMinimumPathOfTreeDFS(root *TreeNode) int {
	minNum := math.MaxInt64
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		depth++

		if root.Left == nil && root.Right == nil {
			minNum = findMin(depth, minNum)
			return
		}

		dfs(root.Left, depth)
		dfs(root.Right, depth)
		depth--

	}

	dfs(root, 0)

	if minNum == math.MaxInt64 {
		return 0
	}
	return minNum

}

/*
1,

1,  2,  3, 4 , 5
*/

func climbStairDFS(n int) int {

	memo := make(map[int]int)
	var dfs func(leftStarts int) int
	dfs = func(leftStarts int) int {

		if leftStarts < 0 {
			return 0
		}
		if leftStarts == 0 {
			return 1
		}

		if _, ok := memo[leftStarts]; ok {
			return memo[leftStarts]
		}

		memo[leftStarts] = dfs(leftStarts-1) + dfs(leftStarts-2)

		return memo[leftStarts]

	}

	memo[n] = dfs(n)
	return memo[n]

}

/*
DP, DFS
base case, house index is equal to greater than len
max = findMax(n + dfs(n+2) , dfs(n+1))

return max
*/
func houseRober2nd(nums []int) int {

	memo := make(map[int]int)
	var dp func(houseIndex int) int

	// max := math.MinInt32
	dp = func(houseIndex int) int {
		if houseIndex >= len(nums) {
			return 0
		}

		if _, ok := memo[houseIndex]; ok {
			return memo[houseIndex]
		}

		rob := nums[houseIndex] + dp(houseIndex+2)
		skip := dp(houseIndex + 1)

		max := findMax(rob, skip)
		memo[houseIndex] = max
		return memo[houseIndex]

	}

	return dp(0)

}

func houseRober2ndCircle(nums []int) int {

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
		return memo[index]

	}

	if len(nums) > 2 {
		one := dfs(0, nums[:len(nums)-1])
		memo = make(map[int]int)
		two := dfs(0, nums[1:])
		max := findMax(one, two)
		return max

	} else {
		return dfs(0, nums)
	}

}

func coinChangeBFS(coins []int, amount int) int {
	q := make([]int, 0)
	q = append(q, amount)
	numOfCoins := 0

	for len(q) > 0 {
		ql := len(q)
		numOfCoins++

		for i := 0; i < ql; i++ {

			balance := q[0]
			q = q[1:]
			for _, coin := range coins {

				change := balance - coin
				if change == 0 {
					return numOfCoins
				} else if change > 0 {
					q = append(q, change)
				} else {
					continue
				}

			}

		}
	}

	return -1

}

/*
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
*/
func longestPalindrome2pointer(s string) string {
	sLen := len(s)
	var maxStr string
	maxLength := 1
	p1, p2, p3 := 0, 0, 0

	for p1 < sLen {

		p2 = p1
		p3 = p1

		for p2 >= 0 && p3 < sLen {
			if s[p2] != s[p3] {
				break
			}

			if (p3 - p2) > maxLength {
				maxLength = p3 - p2
				maxStr = s[p2:p3]
			}

			p2--
			p3++
		}

		p1++

	}

	return maxStr

}

/*
Example 1:

Input: nums = [9,1,4,2,3,3,7]

Output: 4
Explanation: The longest increasing subsequence is [1,2,3,7], which has a length of 4.

Example 2:

Input: nums = [0,3,1,3,2,3]

n=0, 1, 9
n=1, 1  1
n=2, 2 4
n=3 , 2 2
n=4, 3
n=5, 3,
n=6,4

4, 10, 4, 3, 8, 9

0 1
1 2
2 1
3 1
4 1
5 3


e.g,  n =1 and n =2 is known, find out n = 3
max
if nums[n] > nums[n-1] {
	fn(n) = dp[n-1] + 1
	max = fn(n)
} else {
	fn(n) = dp[n-1]
}


init dp[0-n] to  1

dp[0] = 1

for i [1:n] {

	for j =0;i<i
	if nums[i] > nums[j] {
		dp[i] = max(dp[j]+1, dp[i])
	}

}
*/

func lisOwn(nums []int) int {

	dp := make([]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(dp); i++ {

		for j := 0; j < i; j++ {

			if nums[i] > nums[j] {
				dp[i] = findMax(dp[j]+1, dp[i])
			}

		}
	}

	res := 0
	for _, val := range dp {
		res = findMax(res, val)

	}

	return res

}

func longestIncreasingDP(nums []int) int {
	lenS := len(nums)
	dp := make([]int, lenS)

	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 0; i < lenS; i++ {
		for j := 0; j < i; j++ {

			if nums[i] > nums[j] {
				dp[i] = findMax(dp[i], dp[j]+1)
			}

		}
	}

	res := 0
	for _, val := range dp {
		res = findMax(res, val)
	}

	return res

}

/*
Input: s = "12"

Output: 2

Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
12
1, 2
12

base case:
if str starts with 0{
	return 0
}

if str[index] is valid{
	return 1
}




*/

func decodeways(s string) int {

	originalLen := len(s)

	memo := make(map[int]int)

	var dfs func(s string, index int) int
	dfs = func(s string, index int) int {

		if _, ok := memo[index]; ok {
			return memo[index]
		}
		if index >= originalLen {
			return 1
		}

		if s[index] == '0' {
			return 0
		}

		ways := dfs(s, index+1)

		if index+1 < originalLen {
			if (s[index] == '1') || (s[index] == '2' && s[index+1] < '7') {
				ways += dfs(s, index+2)
			}
		}
		memo[index] = ways

		return memo[index]

	}

	tt := dfs(s, 0)
	return tt

}

/*
 */
func uniquePathsBacktracking(m int, n int) int {

	memo := make(map[string]int)
	var dfs func(m int, n int) int
	dfs = func(m int, n int) int {

		stringified := fmt.Sprintf("%d_%d", m, n)
		if memo[stringified] != 0 {
			return memo[stringified]
		}
		if (n == 0) || (m == 0) {
			return 1
		}

		res := dfs(m-1, n) + dfs(m, n-1)
		stringified = fmt.Sprintf("%d_%d", m, n)
		memo[stringified] = res
		return memo[stringified]

	}

	rt := dfs(m-1, n-1)
	return rt

}

/*
Input: nums = [1,2,-3,4]

Output: 4
Example 2:

Input: nums = [-2,-1]

Output: 2
*/
func maxProductBF(nums []int) int {

	max := nums[0]

	for i := 0; i < len(nums); i++ {
		max_temp := nums[i]
		for j := i + 1; j < len(nums); j++ {
			max_temp = findMax(max_temp, nums[i]*nums[j])
			if max_temp == max_temp || nums[i]*nums[j] < max_temp {
				break
			}
		}
		max = findMax(max, max_temp)

	}

	return max

}

/*

1 2 -3 4


*/

func maxProductDPTopDown(nums []int) int {

	max := math.MinInt32
	var dfs func(index int, nums []int, sum int)
	dfs = func(index int, nums []int, sum int) {

		if index == len(nums) {
			return
		}

		if sum > max {
			max = sum
		}

		for i := index + 1; i < len(nums); i++ {
			dfs(i, nums, nums[i]*sum)
		}

	}

	dfs(0, nums, nums[0])

	return max

}

/*
Example 1:

Input: s = "neetcode", wordDict = ["neet","code"]

Output: true
Explanation: Return true because "neetcode" can be split into "neet" and "code".

Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen","ape"]

Output: true

func dfs(index int)

	if index == len(s) {
		found = true
		return
	}

	for word in wordList {
		n = len(word)
		if len(word) < len(s) && s[index:n] == word
			dfs(index + n)
		}
	}

dfs(0)
*/
func wordBreak(s string, wordDict []string) bool {

	memo := make(map[int]int)
	var dfs func(index int) bool
	dfs = func(index int) bool {

		if _, ok := memo[index]; ok {
			if memo[index] == 1 {
				return true
			} else {
				return false
			}

		}

		if index == len(s) {
			return true
		}

		for _, val := range wordDict {
			wordLen := len(val)

			if index+wordLen <= len(s) && s[index:index+wordLen] == val {

				if dfs(index + wordLen) {
					memo[index] = 1
					return true
				}

			}

		}
		memo[index] = 0

		return false

	}

	return dfs(0)

}

/*
Example 1:

Input: nums = [9,1,4,2,3,3,7]

0,3,1,3,2,3

Output: 4,  1 2 3 7
dp = make([]int, len(nums))

0 1
1 1
2 2
3 2
4 3
5 3
6 4

dp[i] = dp[i-1] + 1 if dp[i] > dp[0->i-1]

	for i =0:n {
	   for j = 0->i {
		  if dp[i] > dp[j] {
			 dp[i] = max(dp[i], dp[j] + 1)
		  } else {
			 dp[i] = max(dp[i], dp[j])
		  }
	   }
	}

Explanation: The longest increasing subsequence is [1,2,3,7], which has a length of 4.

Example 2:

Input: nums = [0,3,1,3,2,3]
4 10 4

Output: 4
*/
func lisDP(nums []int) int {

	dp := make([]int, len(nums))

	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = findMax(dp[i], dp[j]+1)
			}

		}

	}

	return dp[len(dp)-1]

}

/*
Input: paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
Output: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]

contecnt1 : [filePath]
contecnt2 : [filePath]

for path in pathts {
	folder = getFolder(path)

	for file in files {
		content = getContent(file)
		map[content] = append(map[content], folder + file)
	}
}


*/

func findDuplicateInFileSystemDFS(root *File) {

	var dfs func(root *File)
	dudupe := make(map[uint64][]string)
	dfs = func(root *File) {

		if root.child == nil {
			if dudupe[root.checksum] == nil {
				dudupe[root.checksum] = make([]string, 0)
			}
			dudupe[root.checksum] = append(dudupe[root.checksum], root.name)

		}

		for _, child := range root.child {
			dfs(child)
		}

	}

}

func findDuplicateInFileSystemBFS(root *File) {

	q := make([]*File, 0)
	q = append(q, root)
	dudupe := make(map[uint64][]string)

	for len(q) > 0 {

		l := len(q)

		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]

			if node.child == nil {
				if dudupe[root.checksum] == nil {
					dudupe[root.checksum] = make([]string, 0)
				}
				dudupe[root.checksum] = append(dudupe[root.checksum], root.name)
			} else {
				for _, child := range root.child {
					q = append(q, child)
				}

			}

		}

	}
}

/*
redo trees
*/

func postorder(root *Node) []int {

	res := make([]int, 0)
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}

		for _, child := range root.Children {
			dfs(child)

		}
		res = append(res, root.Val)

	}

	dfs(root)
	return res

}

func invertBinaryTree2nd(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {

		if root == nil {
			return
		}

		temp := root.Left
		root.Left = root.Right
		root.Right = temp

		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)

	return root

}

func invertBinaryTreeBFS2(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 0)

	q = append(q, root)

	for len(q) != 0 {
		// ql := len(q)

		// for range ql {
		node := q[0]
		if node == nil {
			continue
		}
		temp := node.Left
		node.Left = node.Right
		node.Right = temp

		q = q[1:]

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		// }
	}
	return root

}

func maxDepth3PreOrder(root *TreeNode) int {
	var dfs func(root *TreeNode, depth int)
	maxDepth := 1

	if root == nil {
		return 0
	}

	dfs = func(root *TreeNode, depth int) {

		if root == nil {
			return
		}

		if depth > maxDepth {
			maxDepth = depth
		}

		if root.Left != nil {
			dfs(root.Left, depth+1)
		}

		if root.Right != nil {
			dfs(root.Right, depth+1)
		}

	}

	dfs(root, 1)
	return maxDepth
}

func maxDepth3PostOrder(root *TreeNode) int {
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		left := dfs(root.Left) + 1
		right := dfs(root.Right) + 1

		return findMax(left, right)

	}

	return dfs(root)
}

func diameterOfBinaryTree2(root *TreeNode) int {

	var dfs func(root *TreeNode) int
	maxLenth := 0

	dfs = func(root *TreeNode) int {

		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if left+right > maxLenth {
			maxLenth = left + right
		}
		return findMax(left, right) + 1

	}

	dfs(root)
	return maxLenth

}

func diameterOfBinaryTree2Preorder(root *TreeNode) int {

	var dfs func(root *TreeNode, depth int)
	maxLenthL := 0
	maxLenthR := 0
	max := 0

	dfs = func(root *TreeNode, depth int) {

		if root == nil {
			return
		}

		maxLenthL = findMax(maxLenthL, depth)
		if root.Left != nil {

			dfs(root.Left, depth+1)
		}

		maxLenthR = findMax(maxLenthR, depth)

		if root.Right != nil {
			dfs(root.Left, depth+1)
		}

		max = findMax(max, maxLenthL+maxLenthR)

	}

	dfs(root, 0)
	return max

}

func isSameTree2Pre(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree2Pre(p.Left, q.Left) && isSameTree2Pre(p.Right, q.Right)

}

func lcaDFS(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lcaDFS(root.Left, p, q)
	right := lcaDFS(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return nil

}

func isBalacnedTree(root *TreeNode) bool {

	ret := true

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {

		if ret == false {
			return -1
		}
		if root == nil {
			return 0
		}

		left := dfs(root.Left) + 1
		right := dfs(root.Right) + 1

		if absolute(left, right) > 1 {
			ret = false
		}

		return findMax(left, right)

	}

	if dfs(root) == -1 {
		return false
	}

	return ret

}

func isSameTree3(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || (p.Val != p.Val) {
		return false
	}

	return isSameTree3(p.Left, q.Left) && isSameTree3(p.Right, q.Right)
}

func subTreeOfAnotherTree3(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil && subRoot == nil {
		return true
	}

	if isSameTree3(root, subRoot) {
		return true
	}

	return subTreeOfAnotherTree3(root.Left, subRoot) || subTreeOfAnotherTree3(root.Right, subRoot)

}

/*


dfs(low, high)

mid = low + (high-low) / 2
if root == nil {
	return TreeNode(nums[mid])

}

root.Left = traverse(low, mid-1, root)
root.Right = traverse(mid+1, high, root])



dfs(0, len(nums)/2, root)
*/

func convertSortedArraytoBST(nums []int) *TreeNode {

	var dfs func(nums []int) *TreeNode

	dfs = func(nums []int) *TreeNode {

		if len(nums) == 0 {
			return nil
		}

		mid := len(nums) / 2

		root := &internal.TreeNode{Val: nums[mid]}

		root.Left = dfs(nums[:mid])
		root.Right = dfs(nums[mid+1:])
		return root

	}

	return dfs(nums)

}

/*

if root and root2 are not nil {
	root.Val += root2.val
}


*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var traverse func(root1 *TreeNode, root2 *TreeNode) *TreeNode

	traverse = func(root1 *TreeNode, root2 *TreeNode) *TreeNode {
		if root1 == nil && root2 == nil {
			return nil
		}

		var root *TreeNode
		if root1 != nil && root2 != nil {
			root = &internal.TreeNode{Val: root1.Val + root2.Val}
			root.Left = traverse(root1.Left, root2.Left)
			root.Right = traverse(root1.Right, root2.Right)

		} else if root1 != nil {
			root = &internal.TreeNode{Val: root1.Val}
			root.Left = traverse(root1.Left, nil)
			root.Right = traverse(root1.Right, nil)

		} else {
			root = &internal.TreeNode{Val: root2.Val}
			root.Left = traverse(nil, root2.Left)
			root.Right = traverse(nil, root2.Right)
		}

		return root

	}

	return traverse(root1, root2)

}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if targetSum-root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

}

func rangeSumBST(root *TreeNode, low int, high int) int {

	res := 0
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		if root.Val >= low && root.Val <= high {
			res += root.Val

		}
		traverse(root.Right)
	}

	traverse(root)

	return res

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	r1 := make([]int, 0)
	r2 := make([]int, 0)
	var traverse func(root *TreeNode, r *[]int)

	traverse = func(root *TreeNode, r *[]int) {
		if root == nil {
			return
		}

		traverse(root.Left, r)
		traverse(root.Right, r)
		*r = append(*r, root.Val)

	}

	traverse(root1, &r1)
	traverse(root2, &r2)

	if len(r1) != len(r2) {
		return false
	}

	for i, _ := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true

}

/*
Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
Output: [50,20,80,15,17,19]
parent, child, isLeft

1. find the root node
2. flat the node and make tree structure, then connect them together
3. return the root node


*/

func createBinaryTree(descriptions [][]int) *TreeNode {

	table := make(map[int][]int, 0)

	for _, val := range descriptions {
		if _, ok := table[val[0]]; !ok {
			table[val[0]] = make([]int, 2)
		}
		if val[2] == 1 {
			table[val[0]][0] = val[1]
		} else {
			table[val[0]][1] = val[1]
		}
	}

	rootV := findRoot(table)

	var dfs func(num int) *TreeNode

	dfs = func(num int) *TreeNode {

		root := &internal.TreeNode{Val: num}

		if _, ok := table[num]; ok {
			// return nil

			if table[num][0] != 0 {
				root.Left = dfs(table[num][0])
			}

			if table[num][1] != 0 {
				root.Right = dfs(table[num][1])
			}
		}

		return root

	}
	t := dfs(rootV)

	return t
}

func findRoot(table map[int][]int) int {

	for k, _ := range table {
		flag := 0
		for _, v1 := range table {

			for _, val := range v1 {
				if k == val {
					flag = 1
				}
			}

		}

		if flag == 0 {
			return k
		}
	}

	return 0
}

/*

if root == nil {
	return nil
}

if root == p || root == q {
	return root
}

left := postTraver(root.Left, p, q)
right := postTraver(root.Right, p, q)

if left != nil && right !=nil {
	return root
}

if left != nil {
	return left
}

if right != nil {
	return right
}



return nil




*/

func lowestCommonAncestor3(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)

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

func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &internal.TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = insertBST(root.Right, val)
	}

	return root

}

/*

if node does not have any children, return nil
if ndoe has one child, replace the node with the child
if node has 2 children, pick the smallest from left and replace it with deleted node,
then delete the picked minimum value, recursively







*/

func deletenodeBST(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left != nil && root.Right != nil {

			minVal := bstGetMinValue(root.Left)
			root.Val = minVal
			return deletenodeBST(root.Left, minVal)

		}

		if root.Left != nil {
			temp := root
			root = root.Left
			return temp

		}

		if root.Right != nil {
			temp := root
			root = root.Right
			return temp
		}
	}

	if key < root.Val {
		deletenodeBST(root.Left, key)
	} else {
		deletenodeBST(root.Right, key)
	}

	return root

}

func bstGetMinValue(root *TreeNode) int {

	if root.Left == nil {
		return root.Val
	}

	return bstGetMinValue(root.Left)

}

func btLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	q := make([]*TreeNode, 0)

	q = append(q, root)

	for len(q) != 0 {
		temp := make([]int, 0)

		l := len(q)

		for i := 0; i < l; i++ {
			node := q[0]

			q = q[1:]
			temp = append(temp, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		ret = append(ret, temp)

	}

	return ret
}

func btRightView(root *TreeNode) []int {

	ret := make([]int, 0)

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level == len(ret) {
			ret = append(ret, root.Val)
		}

		dfs(root.Right, level+1)

		dfs(root.Left, level+1)
	}

	dfs(root, 0)

	return ret

}

/*
maintain a max number at every traverse


*/

func goodNodesTree2(root *TreeNode) int {

	res := 0
	var dfs func(root *TreeNode, max int)
	max := math.MinInt64

	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}

		if root.Val > max {
			res++
			max = root.Val
		}

		dfs(root.Left, max)
		dfs(root.Right, max)

	}

	dfs(root, max)

	return res

}

func isValidBST3(root *TreeNode) bool {

	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Val <= root.Left.Val {

		return false

	}

	if root.Right != nil && root.Val >= root.Right.Val {

		return false

	}

	return isValidBST3(root.Left) && isValidBST3(root.Right)

}

func isValidBSTPostOrder(root *TreeNode) bool {

	max := math.MinInt64
	res := true
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)

		if root.Val <= max {
			res = false
		} else {
			max = root.Val
		}
		dfs(root.Right)

	}

	dfs(root)
	return res
}

func kthSmallest(root *TreeNode, k int) int {

	res := 0
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {

		if root == nil {
			return
		}

		dfs(root.Left)
		if k == 1 {
			res = root.Val

		}
		k--

		dfs(root.Right)

	}

	dfs(root)
	return res
}

/*
 preorder: 1 2 5 6 3 7 4
 inorder:  5 2 6 1 7 3 4

 get first index node from preorder as parent
 indexOfParent = len(inorder) / 2 in inorder array.
 build a node
 node.Left = dfs(inorder[:indexOfParent-1])
 node.Right = dfs(inorder[indexOfParent+1:])
 return node


 preorder: 1 2 5 6 3 7 4
 inorder:  5 2 6 1 7 3 4


 pre oreder, inoreder pattern


*/

func constructBTPreInorder(preorder []int, inorder []int) *TreeNode {

	var dfs func(preorder []int, inorder []int) *TreeNode

	dfs = func(preorder []int, inorder []int) *TreeNode {

		if len(preorder) == 0 {
			return nil
		}

		rootVal := preorder[0]

		indexOfRoot := findRootIndex(inorder, rootVal)

		root := &internal.TreeNode{Val: rootVal}

		root.Left = dfs(preorder[1:indexOfRoot+1], inorder[:indexOfRoot])
		root.Right = dfs(preorder[indexOfRoot+1:], inorder[indexOfRoot+1:])

		return root
	}

	root := dfs(preorder, inorder)

	return root

}

func findRootIndex(arr []int, rootVal int) int {

	for i, _ := range arr {
		if rootVal == arr[i] {
			return i
		}
	}

	return -1
}

/*
	if root == nil {
		return nil
	}

root.left = dfs(root.left)
root.right = dfs(root.right)

	if root.val == target && (root.left == nil && root.right == nil){
		return nil
	}

return root
*/
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Val == target && (root.Left == nil && root.Right == nil) {
		return nil
	}

	return root

}

/*
appraoch,
DFS, iterate every island spot, with below base conditions.
The coordiante is in boundary , e.g, in the water close to island or if i,j is island AND i-1 == -1

*/

func islandPerimeter2(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	p := 0

	visted := make(map[[2]int]bool, 0)

	var dfs func(i int, j int)

	dfs = func(i int, j int) {

		if i >= m || j >= n || i < 0 || j < 0 || grid[i][j] == 0 {
			p++
			return
		}

		visitedSpot := [2]int{i, j}
		if _, ok := visted[visitedSpot]; ok {
			return
		}

		visted[visitedSpot] = true

		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == 1 {
				dfs(i, j)
			}

		}
	}

	return p

}

func checkIfLexicographical(words []string, order string) bool {

	if len(words) <= 1 {
		return true
	}

	index := 0
	for i := 1; i < len(words); i++ {
		preWord := words[i-1]
		currWord := words[i]

	}

}

func isAlienSorted(words []string, order string) bool {

}
