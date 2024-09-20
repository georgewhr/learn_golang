package main

import (
	"fmt"
	"george/internal"
)

type ListNode = internal.ListNode
type List = internal.LinkedList

func main() {
	fmt.Println("Hello, world.")
	b := []int{2, 7, 11, 13}
	nums := twoSum(b, 9)
	fmt.Println(nums)

	// a := &b
	// fmt.Println("asdasd %d", (*a)[0])

	myList := List{}
	myList.Insert(9)
	myList.Insert(9)
	myList.Insert(9)

	myList1 := List{}
	myList1.Insert(9)
	myList1.Insert(9)
	myList1.Insert(9)

	myList.PrintList()
	myList1.PrintList()

	myList2 := List{}
	myList2.Head = add2List(myList.Head, myList1.Head)
	myList2.PrintList()

	// myHashMap := internal.InitHashMap()

	// fmt.Println(myHashMap.BackingArr[0])

}

func printLinkedList(head *internal.ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

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
