package internal

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (n *LinkedList) Insert(data int) {
	var newNode = &ListNode{Val: data}
	if n == nil {
		n.Head = newNode
	} else {
		newNode.Next = n.Head
		n.Head = newNode
	}
}

func (n *LinkedList) PrintList() {
	temp := n.Head
	for temp != nil {
		fmt.Printf("%d, ", temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
}
