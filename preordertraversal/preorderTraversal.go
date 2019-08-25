package preordertraversal

import "github.com/dsago/types"

func PreorderTraversal(root *types.TreeNode) []int {
	var slice []int
	run(root, &slice)
	return slice
}

func run(root *types.TreeNode, accumulator *[]int) {
	if root == nil {
		return
	}

	*accumulator = append(*accumulator, root.Val)
	run(root.Left, accumulator)
	run(root.Right, accumulator)
}
