package main

import (
	"fmt"

	"github.com/dsago/breadthfirsttraversal"
	"github.com/dsago/inordertraversal"
	"github.com/dsago/postordertraversal"
	"github.com/dsago/preordertraversal"
	"github.com/dsago/types"
)

func main() {
	var leftNode = types.TreeNode{2, nil, nil}
	var rightNode = types.TreeNode{3, nil, nil}
	var node = types.TreeNode{1, &leftNode, &rightNode}

	var inorderTraversalResults = inordertraversal.InorderTraversal(&node)
	var preorderTraversalResults = preordertraversal.PreorderTraversal(&node)
	var postorderTraversalResults = postordertraversal.PostorderTraversal(&node)
	var breadthfirstTraversalResults = breadthfirsttraversal.BreadthFirstTraversal(&node)
	fmt.Println(inorderTraversalResults)
	fmt.Println(preorderTraversalResults)
	fmt.Println(postorderTraversalResults)
	fmt.Println(breadthfirstTraversalResults)
}
