package arrayqueue

type ArrayQueue struct {
	queueFront int
	queueBack  int
	array      []interface{}
	queueSize  int
}

func New() *ArrayQueue {
	array := make([]interface{}, 4)
	return &ArrayQueue{
		queueFront: 0,
		queueBack:  1,
		array:      array,
		queueSize:  0,
	}
}

func (q *ArrayQueue) Empty() bool {
	// return q.queueFront == q.queueBack
	return q.queueSize == 0
}

func (q *ArrayQueue) Size() int {
	return q.queueSize
}

func (q *ArrayQueue) Front() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.array[q.queueFront+1], true
}

func (q *ArrayQueue) Back() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.array[q.queueBack], true
}

func (q *ArrayQueue) Pop() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	front := (q.queueFront + 1) % len(q.array)
	val := q.array[front]
	q.array[front] = nil
	q.queueFront = front
	q.queueSize--
	return val, true
}

func (q *ArrayQueue) Push(val interface{}) bool {
	if q.Size() == len(q.array)-1 {
		q.array = append(q.array, make([]interface{}, len(q.array)))
	}
	back := (q.queueBack + 1) % len(q.array)
	q.array[back] = val
	q.queueBack = back
	q.queueSize++
	return true
}
