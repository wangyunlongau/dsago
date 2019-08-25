package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var leftNode = TreeNode{2, nil, nil}
	var rightNode = TreeNode{3, nil, nil}
	var node = TreeNode{1, &leftNode, &rightNode}
	var results = inorderTraversal(&node)
	fmt.Println(results)
}

func inorderTraversal(root *TreeNode) []int {
	var slice []int
	return run(root, slice)
}

func run(root *TreeNode, accumulator []int) []int {
	if root == nil {
		return accumulator
	}

	accumulator = append(accumulator, run(root.Left, accumulator)...)
	accumulator = append(accumulator, root.Val)
	accumulator = append(accumulator, run(root.Right, []int{})...)
	return accumulator
}
