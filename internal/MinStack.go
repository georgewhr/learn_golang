package internal

// type Stack = internal.Stack
/*
MinStack API
Push, int
Pop, nil, int
Top,nil. int
GetMin, int

[]int{val, minVal}



*/

type MinStack struct {
	myStack Stack
	cap     int
	size    int
	minVal  int
}

func Construcutor(size int) *MinStack {
	return &MinStack{myStack: *InitStack(size), cap: size}
}

func (s *MinStack) Push(val int) {
	if s.size == 0 {
		s.myStack.Push([]int{val, val})
		s.minVal = val
	} else {
		if val < s.minVal {
			s.minVal = val
		}
		s.myStack.Push([]int{val, s.minVal})
	}

}

// func (s *MinStack) Pop() int {
// 	// return s.myStack.Pop()[0]

// }

// func (s *MinStack) Top() int {

// }

// func (s *MinStack) GetMin() int {

// }
