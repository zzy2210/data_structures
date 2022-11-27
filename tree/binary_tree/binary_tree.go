package binarytree

import (
	queue "data_structures/queue/linked_queue"
	"fmt"
)

type BinaryTree struct {
	size int
	Node *Node
}

type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

func (t *BinaryTree) Empty() bool {
	return t.size == 0
}

func (t *BinaryTree) PreOrder() {
	preOrder(t.Node)
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Println(node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func (t *BinaryTree) InOrder() {
	inOrder(t.Node)
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Println(node)
		inOrder(node.Right)
	}
}

func (t *BinaryTree) PostOrder() {
	postOrder(t.Node)
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Println(node.Val)
	}
}

func (t *BinaryTree) LevelOrder() {
	queue := queue.LinkedQueue{}
	node := t.Node

	for node != nil {
		fmt.Println(node.Val)
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
		queueNode, _ := queue.Pop()
		node = queueNode.(*Node)
	}
}
