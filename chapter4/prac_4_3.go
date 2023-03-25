package ch4

import "github.com/asadzeynal/ctci_practice/util"

type LLNodeOfNodes struct {
	data *util.TreeNode
	next *LLNodeOfNodes
}

func listOfDepths(root *util.TreeNode) *[]*LLNodeOfNodes {
	listOfLists := make([]*LLNodeOfNodes, 0)
	listRecurse(root, &listOfLists, 0)

	return &listOfLists
}

func listRecurse(node *util.TreeNode, listOfLists *[]*LLNodeOfNodes, level int) {
	if node == nil {
		return
	}
	llNode := &LLNodeOfNodes{data: node}
	if level == len(*listOfLists) {
		*listOfLists = append(*listOfLists, llNode)
	} else {
		llNode.next = (*listOfLists)[level]
		(*listOfLists)[level] = llNode
	}

	listRecurse(node.Left, listOfLists, level+1)
	listRecurse(node.Right, listOfLists, level+1)
}
