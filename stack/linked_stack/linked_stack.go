package linkedstack

//  以头节点为栈顶

type LinkedStack struct {
	ChainNode *chainNode
	StackSize int
}

type chainNode struct {
	Val  interface{}
	Next *chainNode
}

func (s *LinkedStack) Empty() bool {
	return s.StackSize == 0
}

func (s *LinkedStack) Size() int {
	return s.StackSize
}

func (s *LinkedStack) Top() (val interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.ChainNode.Val, true
}

func (s *LinkedStack) Pop() (val interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	val = s.ChainNode.Val
	s.ChainNode = s.ChainNode.Next
	return val, true
}

func (s *LinkedStack) Push(val interface{}) (ok bool) {
	s.ChainNode.Next = &chainNode{
		Val:  val,
		Next: s.ChainNode,
	}
	return true
}
