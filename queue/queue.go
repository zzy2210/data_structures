package queue

type Queue interface {
	Empty() bool
	Size() (int, bool)
	// 返回队首值
	Front() (interface{}, bool)
	// 返回队尾值
	Back() (interface{}, bool)
	// 返回并删除队首元素
	Pop() (interface{}, bool)
	// 将元素插入队尾
	Push(interface{}) bool
}
