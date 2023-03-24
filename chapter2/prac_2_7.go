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
