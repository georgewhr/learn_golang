package main

import (
	"container/heap"
	"fmt"
	"george/internal"
)

type ListNode = internal.ListNode
type List = internal.LinkedList
type PriorityQueue = internal.PriorityQueue
type Stack = internal.Stack

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
	myList1.Insert(5)
	myList1.Insert(3)
	myList1.Insert(1)

	myList2 := List{}
	myList2.Insert(6)
	myList2.Insert(4)
	myList2.Insert(2)
	newHead := mergedSortedList(myList1.Head, myList2.Head)
	fmt.Println(newHead)

	// myList.PrintList()
	// myList1.PrintList()

	// myList2 := List{}
	// myList2.Head = add2List(myList.Head, myList1.Head)
	// myList2.PrintList()

	// fmt.Println("length is ", lengthOfLongestSubstring("pwwkew"))

	// myHashMap := internal.InitHashMap()

	// fmt.Println(myHashMap.BackingArr[0])

	// myStack := Stack{}
	// myStack.Push('a')
	// myStack.Push('b')
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

func removeElement(nums []int, val int) int {

}
