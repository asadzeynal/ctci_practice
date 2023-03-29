package ch4

import "github.com/asadzeynal/ctci_practice/util"

func commonAncestor(root *util.TreeNode, a *util.TreeNode, b *util.TreeNode) *util.TreeNode {
	hasA := find(root.Left, a)
	hasb := find(root.Left, b)

	if (hasA && !hasb) || (hasb && !hasA) {
		return root
	}

	var next *util.TreeNode
	if hasA {
		next = root.Left
	} else {
		next = root.Right
	}

	return commonAncestor(next, a, b)
}

func find(root *util.TreeNode, node *util.TreeNode) bool {
	if root == nil {
		return false
	}

	if root == node {
		return true
	}

	return find(root.Left, node) || find(root.Right, node)
}
