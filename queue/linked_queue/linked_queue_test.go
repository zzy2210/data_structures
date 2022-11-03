package linkedqueue_test

import (
	linkedqueue "data_structures/queue/linked_queue"
	"testing"
)

func TestPush(t *testing.T) {
	q := linkedqueue.LinkedQueue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if get := q.Size(); get != 3 {
		t.Errorf("Got %v expected %v", get, 3)
	}
	val := 1
	for node := q.QueueFront; node != nil; node = node.Next {
		if node.Val.(int) != val {
			t.Errorf("Got %v expected %v", node.Val, val)
		}
		val++
	}
}

func TestPop(t *testing.T) {
	q := linkedqueue.LinkedQueue{}
	_, ok := q.Pop()
	if ok {
		t.Errorf("Got %v expect %v", ok, false)
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	get, ok := q.Pop()
	if !ok {
		t.Errorf("Got %v expect %v", ok, true)
	}
	if get != 1 {
		t.Errorf("Got %v expect %v", get, 1)
	}
	if get := q.Size(); get != 2 {
		t.Errorf("Got %v expect %v", get, 2)
	}
	val := 2
	for node := q.QueueFront; node != nil; node = node.Next {
		if node.Val.(int) != val {
			t.Errorf("Got %v expected %v", node.Val, val)
		}
		val++
	}
}

func TestFront(t *testing.T) {
	q := linkedqueue.LinkedQueue{}
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
	q := linkedqueue.LinkedQueue{}
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
