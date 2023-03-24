package ch2

import (
	"github.com/asadzeynal/ctci_practice/util"
)

func isLinkedListPalindromeV2(n *util.Node) bool {
	l := util.LinkedListLength(n)
	_, res := recurseV2(n, l)
	return res
}

func recurseV2(head *util.Node, length int) (*util.Node, bool) {
	if length <= 0 {
		return head, true
	}

	if length == 1 {
		return head.Next, true
	}

	node, isPalindrome := recurseV2(head.Next, length-2)

	return node.Next, isPalindrome && head.Data == node.Data
}

// *****************************************************************
func isLinkedListPalindrome(n *util.Node) bool {
	l := util.LinkedListLength(n)
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
