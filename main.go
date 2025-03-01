package main

import (
	"container/heap"
	"fmt"
	"george/internal"
	"math"
)

type ListNode = internal.ListNode
type List = internal.LinkedList
type PriorityQueue = internal.PriorityQueue
type Stack = internal.Stack
type TreeNode = internal.TreeNode

func main() {
	fmt.Println("Hello, world.")

	// test := []string{"cir", "car"}
	test := []string{"fla", "flw"}
	test1 := "abc"
	fmt.Printf("number is %s", test1[0:1])
	output := longestCommonPrefix(test)
	fmt.Printf("number is %f", output)

	myList := List{}
	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)
	myList.Insert(4)
	myList.Insert(5)

	myNode := removeNthNodeFromEndofList(myList.Head, 2)
	fmt.Printf("number is %f", myNode.Val)

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
	myList1.Insert(1)
	myList1.Insert(2)
	myList1.Insert(3)
	myList1.Insert(4)
	myList1.Insert(5)

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

// 1,2,3
// 3,4,5
// 1,2,3,3,4,5, even,middleIndex = (arr[len / 2 - 1] + arr[len / 2 ]) /2

// 1,2,3
// 3,4,5,6
// 1,2,3,3,4,5,6, odd, arr[len / 2]

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
*/

// func threeSum(number []int) [][]int {

// }

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
		if s[i] == '{' {
			myLocalStack.Push('}')

		} else if s[i] == '[' {
			myLocalStack.Push(']')

		} else if s[i] == '(' {
			myLocalStack.Push(')')
		} else {

			myLocalStack.Push(s[i])

			if myLocalStack.Size() > 0 && myLocalStack.Pop().(rune) != rune(s[i]) {
				return false
			}
			if s[i] == '}' || s[i] == ')' || s[i] == ']' {
				myLocalStack.Push(s[i])
			}
		}

	}

	if myLocalStack.Size() != 0 {
		return false
	}

	return true

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

func sumNumbersHelper(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}

	sum = sum*10 + root.Val

	return sumNumbersHelper(root.Left, sum) + sumNumbersHelper(root.Right, sum)

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil {
		return false
	}

	if q == nil {
		return false
	}

	if q.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right

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
func twoSum1(arr []int, target int) {

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
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {

}

func threeSum(nums []int) [][]int {

}
