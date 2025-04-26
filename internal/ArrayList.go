package internal

/*
Design resiable array.
backing array.
INIT_CAP
current_size
should adjust(double the cap) if size == cap

It has method Add, Remove, Size
addAtIndex
deleteAtIndex
addFront
deleteFront
addEnd
deleteEnd
reszie
clear
isEmpty

*/

type ArrayList struct {
	backing_arr []int
	size        int
	init_cap    int
}

func InitiArrayList() *ArrayList {
	return &ArrayList{
		size:     0,
		init_cap: 10,
	}
}

func (a *ArrayList) Add(val int) {

	if a.size == a.init_cap {
		a.adjust()
	}

}

func (a *ArrayList) Get(index int) {

}

func (a *ArrayList) adjust() {
	newArr := make([]int, 2*a.size)

	copy(newArr, a.backing_arr)

	a.backing_arr = newArr

}
