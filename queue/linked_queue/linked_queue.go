package linkedqueue

type LinkedQueue struct {
	// 指向链首
	QueueFront *chainNode
	QueueBack  *chainNode
	QueueSize  int
}

type chainNode struct {
	Val  interface{}
	Next *chainNode
}

func (q *LinkedQueue) Empty() bool {
	/*
		if q.QueueBack == q.QueueBack && q.QueueFront == nil {
			return true
		}
		return false
	*/
	return q.QueueSize == 0
}

func (q *LinkedQueue) Size() int {
	return q.QueueSize
}

func (q *LinkedQueue) Front() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.QueueFront.Val, true
}

func (q *LinkedQueue) Back() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.QueueBack.Val, true
}

func (q *LinkedQueue) Pop() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	val := q.QueueFront.Val
	q.QueueFront = q.QueueFront.Next
	q.QueueSize--
	return val, true
}

func (q *LinkedQueue) Push(val interface{}) bool {
	node := &chainNode{
		Val: val,
	}
	if q.Empty() {
		q.QueueFront = node
	} else {
		q.QueueBack.Next = node
	}
	q.QueueBack = node
	q.QueueSize++
	return true
}
