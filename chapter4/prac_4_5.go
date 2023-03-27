package ch4

import "github.com/asadzeynal/ctci_practice/util"

func validateBst(node *util.TreeNode) (int, bool) {
	var isBstLeft, isBstRight bool = true, true
	if node == nil {
		// -1 represents zero value in this case
		return -1, true
	}
	valLeft, isBstLeft := validateBst(node.Left)
	valRight, isBstRight := validateBst(node.Right)
	if !isBstLeft || !isBstRight {
		return node.Data, false
	}

	return node.Data, (node.Data >= valLeft || valLeft == -1) && (node.Data < valRight || valRight == -1)
}
