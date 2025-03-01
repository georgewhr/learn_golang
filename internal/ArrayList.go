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
	backing_arr [8]interface{}
	size        int
	init_cap    int
}

func InitiArrayList() *ArrayList {
	return &ArrayList{
		size:     0,
		init_cap: 10,
	}
}
