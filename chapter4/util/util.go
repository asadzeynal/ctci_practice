package util

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode, ns int, ch rune) {
	if root == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", ch, root.Data)
	PrintTree(root.Left, ns+2, 'L')
	PrintTree(root.Right, ns+2, 'R')
}

func printRecurse(root *TreeNode, needLn bool) {

}
