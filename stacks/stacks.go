package stacks

type Stacks interface {
	// 判断栈是否为空
	Empty() bool
	// 输出栈大小
	Size() int
	// 返回栈顶元素，但不推出
	Top() (val interface{}, ok bool)
	// 推出栈顶元素
	Pop() (val interface{}, ok bool)
	// 向栈顶压入元素
	Push(val interface{}) (ok bool)
}
