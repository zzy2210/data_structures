package arraystack_test

import (
	arraystack "data_structures/stack/array_stack"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := &arraystack.ArrayStack{}
	if get := stack.Empty(); get != true {
		t.Errorf("Got %v expected %v", get, true)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if get := stack.Size(); get != 3 {
		t.Errorf("Got %v expected %v", get, 3)
	}
	for i, get := range stack.Array {
		if i+1 != get {
			t.Errorf("Got %v expected %v", i+1, get)
		}
	}
}
func TestStackTop(t *testing.T) {
	stack := &arraystack.ArrayStack{}
	_, ok := stack.Top()
	if ok {
		t.Errorf("Got %v expect %v", true, false)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	get, ok := stack.Top()
	if !ok {
		t.Errorf("Got %v expect %v", false, true)
	}
	if get.(int) != 3 {
		t.Errorf("Got %v expect %v", get, 3)
	}
}

func TestStackPop(t *testing.T) {
	stack := &arraystack.ArrayStack{}
	_, ok := stack.Pop()
	if ok {
		t.Errorf("Got %v expect %v", ok, false)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	get, ok := stack.Pop()
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
