package main

import "fmt"
import "github.com/dsago/inordertraversal"
import "github.com/dsago/preordertraversal"
import "github.com/dsago/postordertraversal"
import "github.com/dsago/types"

func main() {
	var leftNode = types.TreeNode{2, nil, nil}
	var rightNode = types.TreeNode{3, nil, nil}
	var node = types.TreeNode{1, &leftNode, &rightNode}
	
	var inorderTraversalResults = inordertraversal.InorderTraversal(&node)
	var preorderTraversalResults = preordertraversal.PreorderTraversal(&node)
	var postorderTraversalResults = postordertraversal.PostorderTraversal(&node)
	fmt.Println(inorderTraversalResults)
	fmt.Println(preorderTraversalResults)
	fmt.Println(postorderTraversalResults)
}
