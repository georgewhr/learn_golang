package internal

/*
Implement a Queue using backing array
arr [1,2,3], use 2 index, front and tail
init: f=0, tail =-1
[1], f, tail ++, f = 0, tail = 0
[1,2], f  = 0, tail = 1
[1,2,3], f = 0, tail = 2
[_,2,3], remove f, f++, , tail = 2


init()
Iniitliate a queue with size

add()
Add item to queue, return 0 add ok, otherwise -1


peek()
Get first item of the queue, if non, return nil

size()
Get size of queue

remove()
Remove queue and return value, otheriwse return nil

readjust()
can readjust the size, incerase size


a few conditions
f and t should be able to loop the cap using %
if f == t, means there is only 1 item, after that we can reset f =0, t = -1
if (t + 1) % cap == f, queue is full

*/

type Queue struct {
	arr   []interface{}
	front int
	tail  int
	cap   int
}

func InitQueue(cap int) *Queue {
	return &Queue{
		arr:   make([]interface{}, 0, cap),
		front: 0,
		tail:  -1,
		cap:   cap,
	}
}

func (q *Queue) Add(item interface{}) int {
	if (q.tail+1)%q.cap == q.front && (q.tail != -1) {
		return -1
	}
	q.arr = append(q.arr, item)
	q.tail = (q.tail + 1) % q.cap
	return 0
}

func (q *Queue) Remove() interface{} {
	if q.arr[q.front] != nil {
		val := q.arr[q.front]

		if q.front == q.tail {
			q.front = 0
			q.tail = -1
		} else {
			q.front = (q.front + 1) % q.cap
		}

		return val
	} else {
		return nil
	}

}

func (q *Queue) GetFront() interface{} {

	if q.arr[q.front] != nil {
		return q.arr[q.front]
	} else {
		return nil
	}

}

func (q *Queue) GetTail() interface{} {
	if q.arr[q.tail] != nil {
		return q.arr[q.tail]
	} else {
		return nil
	}
}

func (q *Queue) GetSize() int {
	return len(q.arr)

}
