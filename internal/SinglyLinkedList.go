package internal

/*
Design a singly linkedList
Init
Add(item)
Delete(item)
*/

type SinglyListNode struct {
	Val  interface{}
	Next *SinglyListNode
}

type SinglyLinkedList struct {
	val  interface{}
	Head *SinglyListNode
	Tail *SinglyListNode
}

func InitSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{Head: nil}
}

func (l *SinglyLinkedList) InsertFront(item interface{}) {

	if l.Head == nil {
		l.Head = &SinglyListNode{Val: item, Next: nil}
		l.Tail = l.Head
	} else {
		newNode := &SinglyListNode{Val: item, Next: l.Head}
		l.Head = newNode
	}

}

func (l *SinglyLinkedList) InsertEnd(item interface{}) {

	if l.Head == nil {
		l.Head = &SinglyListNode{Val: item, Next: nil}
		l.Tail = l.Head
	} else {
		newNode := &SinglyListNode{Val: item, Next: nil}
		l.Tail.Next = newNode
		l.Tail = newNode

	}

}

func (l *SinglyLinkedList) Delete(item interface{}) *SinglyListNode {
	temp := l.Head

	if temp.Val == item {
		l.Head = nil
		return temp
	}

	for temp.Next != nil {
		if temp.Next.Val == item {
			if temp.Next == l.Tail {
				l.Tail = temp
			}
			rtNode := temp.Next
			temp.Next = temp.Next.Next
			return rtNode
		}
		temp = temp.Next
	}
	return nil

}
