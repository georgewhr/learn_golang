package internal

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (n *LinkedList) Insert(data int) {
	var newNode = new(Node)
	newNode.Data = data
	if n == nil {
		n.head = newNode
	} else {
		newNode.Next = n.head
		n.head = newNode
	}
}

func (n *LinkedList) PrintList() {
	temp := n.head
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}
