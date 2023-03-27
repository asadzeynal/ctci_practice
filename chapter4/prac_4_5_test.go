package ch4

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestValidateBstTrue(t *testing.T) {
	nodes := []*util.TreeNode{
		{Data: 0},
		{Data: 1},
		{Data: 2},
		{Data: 3},
		{Data: 4},
		{Data: 5},
	}
	nodes[3].Left = nodes[1]
	nodes[3].Right = nodes[4]

	nodes[1].Left = nodes[0]
	nodes[1].Right = nodes[2]
	nodes[4].Right = nodes[5]

	_, ans := validateBst(nodes[3])

	if !ans {
		t.Errorf("wrong answer got: %v, want: %v", ans, true)
	}
}

func TestValidateBstFalse(t *testing.T) {
	nodes := []*util.TreeNode{
		{Data: 0},
		{Data: 1},
		{Data: 2},
		{Data: 3},
		{Data: 4},
		{Data: 5},
	}
	nodes[3].Left = nodes[4]
	nodes[3].Right = nodes[1]

	nodes[1].Left = nodes[0]
	nodes[1].Right = nodes[2]
	nodes[4].Right = nodes[5]

	_, ans := validateBst(nodes[3])

	if ans {
		t.Errorf("wrong answer got: %v, want: %v", ans, false)
	}
}
