package internal

import "fmt"

type ListNode struct {
	Val      int
	Next     *ListNode
	Previous *ListNode
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

// 1 -> 2 ->3 ->4
// 1 -> 2
func (n *LinkedList) deleteList(data int) {
	temp := n.Head
	if data == temp.Val {
		n.Head = temp.Next
	}

	for temp.Next != nil {
		if data == temp.Next.Val {
			temp.Next = temp.Next.Next
		}
		temp = temp.Next
	}
}
