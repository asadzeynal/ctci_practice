package ch2

import "github.com/asadzeynal/ctci_practice/util"

func partition(n *util.Node, p int) *util.Node {
	headSm := &util.Node{}
	smPointer := headSm
	headGt := &util.Node{}
	gtPointer := headGt

	for n != nil {
		if n.Data < p {
			smPointer.Next = &util.Node{Data: n.Data}
			smPointer = smPointer.Next
		} else {
			gtPointer.Next = &util.Node{Data: n.Data}
			gtPointer = gtPointer.Next
		}
		n = n.Next
	}

	headSm = headSm.Next
	headGt = headGt.Next
	smPointer.Next = headGt
	return headSm
}
