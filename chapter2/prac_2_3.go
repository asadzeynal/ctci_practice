package ch2

import "github.com/asadzeynal/ctci_practice/util"

func deleteMiddleNode(n *util.Node) {
	next := n.Next
	n.Data = next.Data
	n.Next = next.Next
}
