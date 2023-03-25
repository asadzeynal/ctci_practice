package ch4

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestMinimalTree(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root := minimalTree(list)
	util.PrintTree(root, 0, 'M')
	isBinarySearchTreeCorrect(root, t)
}

func isBinarySearchTreeCorrect(node *util.TreeNode, t *testing.T) {
	if node.Left == nil {
		return
	} else {
		if node.Data < node.Left.Data {
			t.Errorf("parent less than left child: %v %v", node.Data, node.Left.Data)
		}
	}
	if node.Right == nil {
		return
	} else {
		if node.Data > node.Right.Data {
			t.Errorf("parent more than right child: %v %v", node.Data, node.Right.Data)
		}
	}

	isBinarySearchTreeCorrect(node.Left, t)
	isBinarySearchTreeCorrect(node.Right, t)
}
