package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    
    var validate func(root *TreeNode) (int, bool)
    validate = func(root *TreeNode) (int, bool) {
        if root == nil {
            return 0, true
        }

        lHeight, ok := validate(root.Left)
        if !ok {
            return 0, false
        }
        rHeight, ok := validate(root.Right) 
        if ! ok {
            return 0, false
        }

        diff := rHeight - lHeight
        if -1 > diff || diff > 1 {
            return 0, false
        }

        return max(lHeight, rHeight) + 1, true
    }

    _, ok := validate(root)
    return ok
}