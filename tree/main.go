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
	var nodeSeven = types.TreeNode{7, nil, nil}
	var nodeSix = types.TreeNode{6, nil, nil}
	var nodeFive = types.TreeNode{5, nil, nil}
	var nodeFour = types.TreeNode{4, nil, nil}
	var nodeThree = types.TreeNode{3, &nodeSix, &nodeSeven}
	var nodeTwo = types.TreeNode{2, &nodeFour, &nodeFive}
	var nodeOne = types.TreeNode{1, &nodeTwo, &nodeThree}

	var inorderTraversalResults = inordertraversal.InorderTraversal(&nodeOne)
	var preorderTraversalResults = preordertraversal.PreorderTraversal(&nodeOne)
	var postorderTraversalResults = postordertraversal.PostorderTraversal(&nodeOne)
	var breadthfirstTraversalResults = breadthfirsttraversal.BreadthFirstTraversal(&nodeOne)
	fmt.Println(inorderTraversalResults)
	fmt.Println(preorderTraversalResults)
	fmt.Println(postorderTraversalResults)
	fmt.Println(breadthfirstTraversalResults)
}
