package list

type List interface {
	Empty() bool
	Size() int
	// 根据索引返回对应元素的值
	Get(index int) (val interface{}, ok bool)
	// 根据输入的值返回查找到的第一个索引
	IndexOf(val interface{}) (index int, ok bool)
	// 删除索引对应的节点
	Remove(index int)
	// 在索引为index的位置上插入一个新的元素
	Insert(index int, val interface{}) bool
}
