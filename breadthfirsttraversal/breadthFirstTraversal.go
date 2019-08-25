package breadthfirsttraversal

import "github.com/dsago/types"

func BreadthFirstTraversal(root *types.TreeNode) [][]int {
	var slice [][]int
	if root == nil {
		return slice
	}

	var queue = []*types.TreeNode{root}
	
	for len(queue) != 0 {
		var temp []int
		var i = 0
		var length = len(queue)

		for i < length {
			var e = queue[0]
			queue = queue[1:]

			temp = append(temp, e.Val)
			if e.Left != nil {
				queue = append(queue, e.Left)
			}
			if e.Right != nil {
				queue = append(queue, e.Right)
			}
			i++
		}
		slice = append(slice, temp)
	}
	return slice
}
