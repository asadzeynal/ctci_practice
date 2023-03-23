package ch2

import (
	"github.com/asadzeynal/ctci_practice/util"
)

func removeDups(head *util.Node) *util.Node {
	prev := head
	if prev == nil {
		return head
	}

	curr := head.Next
	if curr == nil {
		return head
	}

	exists := make(map[int]struct{})
	for curr.Next != nil {
		exists[prev.Data] = struct{}{}
		if _, ok := exists[curr.Data]; ok {
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = prev.Next
	}
	if _, ok := exists[curr.Data]; ok {
		prev.Next = nil
	}
	return head
}
