package arrayqueue_test

import (
	arrayqueue "data_structures/queue/array_queue"
	"testing"
)

func TestPush(t *testing.T) {
	q := arrayqueue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if get := q.Size(); get != 3 {
		t.Errorf("Got %v expected %v", get, 3)
	}
	q.Push(4)
	if get := len(q.Array); get != 8 {
		t.Errorf("Got %v expected %v", get, 8)
	}
	if get := q.QueueBack; get != 4 {
		t.Errorf("Got %v expected %v", get, 4)
	}
	if get := q.Size(); get != 4 {
		t.Errorf("Got %v expected %v", get, 4)
	}
}

func TestPop(t *testing.T) {
	q := arrayqueue.New()
	get, ok := q.Pop()
	if ok {
		t.Errorf("Got %v expect %v", ok, false)
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	get, ok = q.Pop()
	if !ok {
		t.Errorf("Got %v expect %v", ok, true)
	}
	if get != 1 {
		t.Errorf("Got %v expect %v", get, 1)
	}
	if get := q.Size(); get != 2 {
		t.Errorf("Got %v expect %v", get, 2)
	}
	if get := q.QueueFront; get != 1 {
		t.Errorf("Got %v expect %v", get, 1)
	}
}

func TestFront(t *testing.T) {
	q := arrayqueue.New()
	_, ok := q.Front()
	if ok {
		t.Errorf("Got %v expect %v", ok, false)
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	get, ok := q.Front()
	if !ok {
		t.Errorf("Got %v expect %v", ok, true)
	}
	if get.(int) != 1 {
		t.Errorf("Got %v expect %v", get, 1)
	}
}

func TestBack(t *testing.T) {
	q := arrayqueue.New()
	if _, ok := q.Back(); ok {
		t.Errorf("Got %v expect %v", ok, 1)
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	get, ok := q.Back()
	if !ok {
		t.Errorf("Got %v expect %v", ok, true)
	}
	if get.(int) != 3 {
		t.Errorf("Got %v expect %v", get, 3)
	}
}
