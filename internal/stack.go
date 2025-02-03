package internal

import "fmt"

/*
Implement a stack.
Push(), insert into a stack
Pop(), pop out item from stack
Top(), return the latest(top) inserted item
isEmpty(), check if stack is empty
*isFull(), check if stack is full, only used if it's backing array
Size(), return size of stack
Requirements:
1. Stack needs to have above operations
2. The stack should be able to adjust its size when it is full

Implementation approach
1. The item will be interface{} for universal item
2. *has a init function if it's backing array, and return a pointer to the stack
3. Decide backing array or backing slice, if the size of data is fixed, use array, ohterwise use slice
*/

// const INIT_SIZE = 12

// backing slice for the stack
type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {

	if len(*s) == 0 {
		return nil
	} else {
		popItemIndex := len(*s) - 1
		i := (*s)[popItemIndex]
		*s = (*s)[:popItemIndex]
		return i
	}
}

func (s *Stack) Peek() interface{} {

	if len(*s) == 0 {
		return nil
	} else {
		popItemIndex := len(*s) - 1
		return (*s)[popItemIndex]
	}
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) PrintStack() {
	fmt.Printf("Slice: %v\n", *s)
}
