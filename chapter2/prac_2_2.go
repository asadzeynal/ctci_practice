package ch2

import "github.com/asadzeynal/ctci_practice/util"

func kToLastIterative(head *util.Node, k int) int {
	counter := 0
	var kTh *util.Node
	n := head
	for n != nil {
		if counter == k {
			kTh = head
		}
		if counter > k {
			kTh = kTh.Next
		}
		n = n.Next
		counter++
	}

	return kTh.Data
}

func kToLastRecursive(head *util.Node, k int) int {
	res, _ := getKthToLastNode(head, k)
	return res.Data
}

func getKthToLastNode(n *util.Node, k int) (*util.Node, int) {
	if n == nil {
		return n, -1
	}
	next, count := getKthToLastNode(n.Next, k)
	if count >= k {
		return next, count
	}
	return n, count + 1
}
