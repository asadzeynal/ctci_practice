package ch4

import (
	"testing"

	"github.com/asadzeynal/ctci_practice/util"
)

// Not really a good test for now
func TestListOfDepths(t *testing.T) {
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

	listOfLists := listOfDepths(nodes[0])

	// for i, v := range *listOfLists {
	// 	for v != nil {
	// 		fmt.Printf("%v %v, ", i, v.data.Data)
	// 		v = v.next
	// 	}
	// 	fmt.Println()
	// }

	if len(*listOfLists) != 5 {
		t.Errorf("wrong list length, got: %v, want: %v", len(*listOfLists), 5)
	}
}
