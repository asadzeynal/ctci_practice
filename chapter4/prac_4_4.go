package ch4

import (
	"math"

	"github.com/asadzeynal/ctci_practice/util"
)

// Not my implementation
func checkBalanced(node *util.TreeNode) bool {
	return checkRecurse(node) != -2
}

func checkRecurse(node *util.TreeNode) int {
	if node == nil {
		return -1
	}

	leftH := checkRecurse(node.Left)
	if leftH == -2 {
		return -2
	}
	rightH := checkRecurse(node.Right)
	if rightH == -2 {
		return -2
	}

	diff := (leftH - rightH)
	if diff > 1 || diff < -1 {
		return -2
	}

	return int(math.Max(float64(leftH), float64(rightH))) + 1
}
