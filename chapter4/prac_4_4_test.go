package ch4

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestCheckBalanced(t *testing.T) {
	nodes := []*util.TreeNode{
		{Data: 0},
		{Data: 1},
		{Data: 2},
		{Data: 3},
		{Data: 4},
		{Data: 5},
	}
	nodes[0].Left = nodes[1]
	nodes[0].Right = nodes[2]

	nodes[1].Left = nodes[3]
	nodes[3].Left = nodes[4]
	nodes[4].Left = nodes[5]

	ans1 := checkBalanced(nodes[0])

	if ans1 {
		t.Errorf("wrong answer got: %v, want: %v", ans1, false)
	}

	ans2 := checkBalanced(nodes[5])

	if !ans2 {
		t.Errorf("wrong answer got: %v, want: %v", ans2, true)
	}
}
