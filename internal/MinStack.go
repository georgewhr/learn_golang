package internal

import "math"

// type Stack = internal.Stack
/*
MinStack API
Push, int
Pop, nil, int
Top,nil. int
GetMin, int

[]int{val, minVal}



*/

type Items struct {
	data   int
	minVal int
}
type MinStack struct {
	myStack      []Items
	cap          int
	currentIndex int
	size         int
	minVal       int
}

func Construcutor(size int) *MinStack {
	return &MinStack{myStack: []Items{}, cap: size, minVal: math.MaxInt, currentIndex: 0}
}

func (s *MinStack) Push(val int) {
	if s.cap == s.size {
		return
	}

	if val < s.minVal {
		s.minVal = val
	}
	data := Items{data: val, minVal: s.minVal}

	s.myStack[s.currentIndex] = data
	s.currentIndex++
}

func (s *MinStack) Pop() int {
	item := s.myStack[s.currentIndex]
	s.currentIndex--
	s.minVal = s.myStack[s.currentIndex].minVal
	return item.data

}

func (s *MinStack) Top() int {
	item := s.myStack[s.currentIndex]
	return item.data

}

func (s *MinStack) GetMin() int {
	return s.minVal

}
