package list

// TODO : defiend list and list node and more

type Node struct {
	Val  interface{}
	Next *Node
}

type List interface {
	Get(val interface{}) (node *Node, ok bool)
	Remove()
	Add()
}
