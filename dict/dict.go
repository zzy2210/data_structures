package dict

type Dict interface {
	Empty() bool
	Size() int
	// 根据k返回v
	Get(key interface{}) (interface{}, bool)
	// 插入键值对
	Insert(key, val interface{}) bool
	// 根据k删除键值对
	Erase(key interface{}) bool
}
