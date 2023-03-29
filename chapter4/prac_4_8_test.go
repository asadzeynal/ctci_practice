package ch4

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestCommonAncestor(t *testing.T) {
	nodes := []*util.TreeNode{
		{Data: 25}, // 0
		{Data: 20}, // 1
		{Data: 36}, // 2
		{Data: 10}, // 3
		{Data: 22}, // 4
		{Data: 30}, // 5
		{Data: 40}, // 6
		{Data: 5},  // 7
		{Data: 12}, // 8
		{Data: 28}, // 9
		{Data: 38}, // 10
		{Data: 48}, // 11
	}
	nodes[0].Left = nodes[1]
	nodes[0].Right = nodes[2]
	nodes[1].Left = nodes[3]
	nodes[1].Right = nodes[4]
	nodes[2].Left = nodes[5]
	nodes[2].Right = nodes[6]
	nodes[3].Left = nodes[7]
	nodes[3].Right = nodes[8]
	nodes[5].Left = nodes[9]
	nodes[6].Left = nodes[10]
	nodes[6].Right = nodes[11]

	ans := commonAncestor(nodes[0], nodes[11], nodes[9])

	if ans != nodes[2] {
		t.Errorf("Wrong answer, got: %v, want: %v", ans.Data, nodes[2].Data)
	}

	ans = commonAncestor(nodes[0], nodes[7], nodes[4])

	if ans != nodes[1] {
		t.Errorf("Wrong answer, got: %v, want: %v", ans.Data, nodes[2].Data)
	}

	ans = commonAncestor(nodes[0], nodes[1], nodes[2])

	if ans != nodes[0] {
		t.Errorf("Wrong answer, got: %v, want: %v", ans.Data, nodes[0].Data)
	}
}
