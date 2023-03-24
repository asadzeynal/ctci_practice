package ch2

import "github.com/asadzeynal/ctci_practice/util"

func intersect(node1 *util.Node, node2 *util.Node) bool {
	for node1.Next != nil {
		node1 = node1.Next
	}
	for node2.Next != nil {
		node2 = node2.Next
	}

	return node1 == node2
}

func intersectWithReturn(node1 *util.Node, node2 *util.Node) *util.Node {
	l1 := util.LinkedListLength(node1)
	l2 := util.LinkedListLength(node2)

	lenDiff := l1 - l2
	if lenDiff < 0 {
		node2 = skipKNodes(node2, lenDiff)
	}
	if lenDiff > 0 {
		node1 = skipKNodes(node1, lenDiff)
	}

	for node1.Next != nil && node2.Next != nil {
		if node1 == node2 {
			return node1
		}
		node1 = node1.Next
		node2 = node2.Next
	}

	return nil
}

func skipKNodes(node *util.Node, k int) *util.Node {
	var count int
	for count < k {
		node = node.Next
		count++
	}
	return node
}
