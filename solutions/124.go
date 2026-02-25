package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    result := -1001

    var dfs func(root *TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        leftMaxPathSum := max(dfs(root.Left), 0)
        rightMaxPathSum := max(dfs(root.Right), 0)

        result = max(result, leftMaxPathSum + root.Val + rightMaxPathSum)
        return root.Val + max(leftMaxPathSum, rightMaxPathSum)
    }
    dfs(root)
    return result
}