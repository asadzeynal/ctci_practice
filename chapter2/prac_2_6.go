package ch2

import (
	"github.com/asadzeynal/ctci_practice/util"
)

func isLinkedListPalindrome(n *util.Node) bool {
	l := length(n)
	stack := make([]int, l)
	return recurse(n, &stack, 0)
}

func recurse(n *util.Node, s *[]int, i int) bool {
	if n == nil {
		return true
	}
	(*s)[i] = n.Data
	isEqual := recurse(n.Next, s, i+1)
	return isEqual && (*s)[len(*s)-i-1] == n.Data
}

func length(head *util.Node) int {
	var l int
	for head != nil {
		l++
		head = head.Next
	}

	return l
}
