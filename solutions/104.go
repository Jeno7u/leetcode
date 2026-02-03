package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    var recursionCheck func(root *TreeNode, depth int) int
    recursionCheck = func(root *TreeNode, depth int) int {
        if root == nil {
            return depth - 1
        }

        return max(recursionCheck(root.Left, depth + 1),
                    recursionCheck(root.Right, depth + 1))
    }

    return recursionCheck(root, 1)
}