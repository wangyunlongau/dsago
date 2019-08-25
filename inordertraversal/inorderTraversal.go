package inordertraversal

import "github.com/dsago/types"

func InorderTraversal(root *types.TreeNode) []int {
	var slice []int
	return run(root, slice)
}

func run(root *types.TreeNode, accumulator []int) []int {
	if root == nil {
		return accumulator
	}

	accumulator = append(accumulator, run(root.Left, accumulator)...)
	accumulator = append(accumulator, root.Val)
	accumulator = append(accumulator, run(root.Right, []int{})...)
	return accumulator
}
