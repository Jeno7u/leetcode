package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var validate func(root *TreeNode, min, max *int) bool
    validate = func(root *TreeNode, min, max *int) bool {
        if root == nil {
            return true
        }

        if root.Left.Val < *min {
            return false
        }

        if root.Right.Val > *max {
            return false
        }

        return validate(root.Left, min, &root.Val) && validate(root.Right, &root.Val, max)
    }

    return validate(root, nil, nil)
}