package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct 
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * 
 */

func rightSideView(root *TreeNode) []int {
    result := []int{}

    var getView func(node *TreeNode, depth int)
    getView = func(node *TreeNode, depth int) {
        if node == nil {
            return
        }

        if depth > len(result) {
            result = append(result, node.Val)
        }

        getView(node.Right, depth+1)
        getView(node.Left, depth+1)
    }

    getView(root, 1)
    return result
}
