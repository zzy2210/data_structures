package linkedstack_test

import (
	linkedstack "data_structures/stack/linked_stack"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := &linkedstack.LinkedStack{}
	if get := stack.Empty(); get != true {
		t.Errorf("Got %v expected %v", get, true)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if get := stack.Size(); get != 3 {
		t.Errorf("Got %v expected %v", get, 3)
	}
	val := 3
	for node := stack.ChainNode; node != nil; node = node.Next {
		if node.Val.(int) != val {
			t.Errorf("Got %v expected %v", node.Val, val)
		}
		val--
	}
}

func TestStackTop(t *testing.T) {
	stack := &linkedstack.LinkedStack{}
	get, ok := stack.Top()
	if ok {
		t.Errorf("Got %v expect %v", true, false)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	get, ok = stack.Top()
	if !ok {
		t.Errorf("Got %v expect %v", false, true)
	}
	if get.(int) != 3 {
		t.Errorf("Got %v expect %v", get, 3)
	}
}

func TestStackPop(t *testing.T) {
	stack := &linkedstack.LinkedStack{}
	get, ok := stack.Pop()
	if ok {
		t.Errorf("Got %v expect %v", true, false)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	get, ok = stack.Pop()
	if !ok {
		t.Errorf("Got %v expect %v", false, true)
	}
	if get.(int) != 3 {
		t.Errorf("Got %v expect %v", get, 3)
	}
	if get := stack.Size(); get != 2 {
		t.Errorf("Got %v expect %v", get, 2)
	}
}
