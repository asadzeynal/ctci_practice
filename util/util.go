package util

type Node struct {
	Data int
	Next *Node
}

func EqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !EqualSlice(v, b[i]) {
			return false
		}
	}
	return true
}

func LinkedListFromSlice(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	head := &Node{}
	n := head
	for i := range s {
		n.Data = s[i]
		if i == len(s)-1 {
			n.Next = nil
			break
		}
		n.Next = &Node{}
		n = n.Next
	}
	return head
}

func SliceFromLinkedList(n *Node) []int {
	result := make([]int, 0)
	for n != nil {
		result = append(result, n.Data)
		n = n.Next
	}
	return result
}
