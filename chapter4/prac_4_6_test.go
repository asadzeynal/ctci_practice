package ch4

import (
	"fmt"
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

func TestSuccessor(t *testing.T) {
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

	nodes[1].Parent = nodes[0]
	nodes[2].Parent = nodes[0]
	nodes[3].Parent = nodes[1]
	nodes[4].Parent = nodes[1]
	nodes[5].Parent = nodes[2]
	nodes[6].Parent = nodes[2]
	nodes[7].Parent = nodes[3]
	nodes[8].Parent = nodes[3]
	nodes[9].Parent = nodes[5]
	nodes[10].Parent = nodes[6]
	nodes[11].Parent = nodes[6]

	tests := []struct {
		node *util.TreeNode
		want *util.TreeNode
	}{
		{nodes[7], nodes[3]},
		{nodes[3], nodes[8]},
		{nodes[8], nodes[1]},
		{nodes[1], nodes[4]},
		{nodes[4], nodes[0]},
		{nodes[0], nodes[9]},
		{nodes[9], nodes[5]},
		{nodes[5], nodes[2]},
		{nodes[2], nodes[10]},
		{nodes[10], nodes[6]},
		{nodes[6], nodes[11]},
		{nodes[11], nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.node.Data)
		t.Run(testname, func(t *testing.T) {
			ans := successor(tt.node)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans.Data, tt.want.Data)
			}
		})
	}
}
