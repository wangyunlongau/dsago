package postordertraversal

import "github.com/dsago/types"

func PostorderTraversal(root *types.TreeNode) []int {
	var slice []int
	run(root, &slice)
	return slice
}

func run(root *types.TreeNode, accumulator *[]int) {
	if root == nil {
		return
	}
	run(root.Left, accumulator)
	run(root.Right, accumulator)
	*accumulator = append(*accumulator, root.Val)
}
