package ch4

import "github.com/asadzeynal/ctci_practice/util"

func successor(node *util.TreeNode) *util.TreeNode {
	if node.Right != nil {
		curr := node.Right
		for curr.Left != nil {
			curr = curr.Left
		}
		return curr
	}

	for node.Parent != nil && node != node.Parent.Left {
		node = node.Parent
	}
	return node.Parent
}
