package main

import (
	"fmt"
	"george/internal"
)

type Node = internal.Node
type List = internal.LinkedList

func main() {
	fmt.Println("Hello, world.")
	b := []int{2, 7, 11, 13}
	nums := twoSum(b, 9)
	fmt.Println(nums)

	// a := &b
	// fmt.Println("asdasd %d", (*a)[0])

	myList := List{}
	myList.Insert(5)
	myList.Insert(6)
	myList.PrintList()
	// myList.PrintList()

	myHashMap := internal.InitHashMap()

	fmt.Println(myHashMap.BackingArr[0])

}

func printLinkedList(head *internal.Node) {
	for head != nil {
		fmt.Println(head.Data)
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

func add2List(l1 *Node, l2 Node) {

}
