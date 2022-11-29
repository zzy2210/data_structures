package binarytree_test

import (
	binarytree "data_structures/tree/binary_tree"
	"testing"
)

func LevelOrderTest(t *testing.T) {
	rootnode := binarytree.Node{
		Val: 0,
	}
	rootleft := binarytree.Node{
		Val: 1,
	}
	rootrigt := binarytree.Node{
		Val: 2,
	}
	rootnode.Left = &rootleft
	rootnode.Right = &rootrigt

	l2 := binarytree.Node{
		Val: 3,
	}
	r2 := binarytree.Node{
		Val: 4,
	}
	rootnode.Left.Left = &l2
	rootnode.Left.Right = &r2

	tree := binarytree.BinaryTree{
		Node: &rootnode,
	}
	tree.LevelOrder()
}
