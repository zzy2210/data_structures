package singledlinkedlist

type SingleChainNode struct {
	Val  interface{}
	Next *SingleChainNode
}

// h means head
func (h *SingleChainNode) Empty() bool {
	return h.Next == nil
}

func (h *SingleChainNode) Size() int {
	if h.Empty() {
		return 0
	}
	var size int
	for node := h; node.Next != nil; node = node.Next {
		size++
	}
	return size
}

func (h *SingleChainNode) Get(indext int) (val interface{}, ok bool) {
	if h.Empty() || h.Size() < indext {
		return nil, false
	}
	// 头节点序列为0
	node := h.Next
	for i := 1; i < indext; i++ {
		node = node.Next
	}
	return node.Val, true
}

func (h *SingleChainNode) IndexOf(val interface{}) (index int, ok bool) {
	if h.Empty() {
		return -1, false
	}
	// 保证可以遍历到最后一个节点
	for node := h; node != nil; node = node.Next {
		index++
		if node.Val == val {
			return index, true
		}
	}
	return -1, false
}

// 思考了一下，哪怕输入了超出界值的索引，也不会输出错误
// 因为Go的GC机制，所以无需手动删除对应内存，去除索引即可
func (h *SingleChainNode) Remove(index int) {
	if h.Empty() {
		return
	}
	if index == 0 {
		deleteNode := h
		h = deleteNode.Next
	}
	node := h
	for i := 0; i < index-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
}

func (h *SingleChainNode) Insert(index int, val interface{}) bool {
	if index < 0 || index > h.Size() {
		return false
	}
	if index == 0 {
		node := &SingleChainNode{
			Val:  val,
			Next: h,
		}
		h = node
		return true
	}
	node := h
	for i := 0; i < index-1; i++ {
		node = node.Next
	}
	node.Next = &SingleChainNode{
		Val:  val,
		Next: node.Next,
	}
	return true
}
