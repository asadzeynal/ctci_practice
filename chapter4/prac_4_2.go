package ch4

import "github.com/asadzeynal/ctci_practice/chapter4/util"

func minimalTree(list []int) *util.TreeNode {
	if len(list) == 0 {
		return nil
	}
	if len(list) == 1 {
		return &util.TreeNode{Data: list[0]}
	}
	middleIndex := len(list) / 2
	node := &util.TreeNode{Data: list[middleIndex]}
	node.Left = minimalTree(list[:middleIndex])
	node.Right = minimalTree(list[middleIndex+1:])
	return node
}
