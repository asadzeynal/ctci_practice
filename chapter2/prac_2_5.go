package ch2

import (
	"github.com/asadzeynal/ctci_practice/util"
)

func sumListsReverseRecurse(n1 *util.Node, n2 *util.Node) *util.Node {
	return sumReverse(n1, n2, 0)
}

func sumReverse(n1 *util.Node, n2 *util.Node, carry int) *util.Node {
	if n1 == nil && n2 == nil && carry == 0 {
		return nil
	}

	num1, num2 := 0, 0
	var next1, next2 *util.Node
	if n1 != nil {
		num1 = n1.Data
		next1 = n1.Next
	}
	if n2 != nil {
		num2 = n2.Data
		next2 = n2.Next
	}
	sum := num1 + num2 + carry
	if sum >= 10 {
		carry = 1
		sum = sum % 10
	} else {
		carry = 0
	}

	result := &util.Node{Data: sum}

	result.Next = sumReverse(next1, next2, carry)
	return result
}

// *******************************************************************************
func sumListsReverse(n1 *util.Node, n2 *util.Node) *util.Node {
	carry := 0
	var head *util.Node
	var n *util.Node
	for n1 != nil || n2 != nil {
		if n == nil {
			n = &util.Node{}
			head = n
		} else {
			n.Next = &util.Node{}
			n = n.Next
		}

		num1, num2 := 0, 0
		if n1 != nil {
			num1 = n1.Data
			n1 = n1.Next
		}

		if n2 != nil {
			num2 = n2.Data
			n2 = n2.Next
		}

		sum := num1 + num2 + carry
		if sum >= 10 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}
		n.Data = sum
	}

	if carry == 1 {
		n.Next = &util.Node{Data: 1, Next: nil}
	}

	return head
}
