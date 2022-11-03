package arrayqueue

type ArrayQueue struct {
	QueueFront int
	QueueBack  int
	Array      []interface{}
	QueueSize  int
}

func New() *ArrayQueue {
	array := make([]interface{}, 4)
	return &ArrayQueue{
		QueueFront: 0,
		QueueBack:  0,
		Array:      array,
		QueueSize:  0,
	}
}

func (q *ArrayQueue) Empty() bool {
	// return q.queueFront == q.queueBack
	return q.QueueSize == 0
}

func (q *ArrayQueue) Size() int {
	return q.QueueSize
}

func (q *ArrayQueue) Front() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.Array[q.QueueFront+1], true
}

func (q *ArrayQueue) Back() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	return q.Array[q.QueueBack], true
}

func (q *ArrayQueue) Pop() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	front := (q.QueueFront + 1) % len(q.Array)
	val := q.Array[front]
	q.Array[front] = nil
	q.QueueFront = front
	q.QueueSize--
	return val, true
}

func (q *ArrayQueue) Push(val interface{}) bool {
	if q.Size() == len(q.Array)-1 {
		q.Array = append(q.Array, make([]interface{}, len(q.Array))...)
	}
	back := (q.QueueBack + 1) % len(q.Array)
	q.Array[back] = val
	q.QueueBack = back
	q.QueueSize++
	return true
}
