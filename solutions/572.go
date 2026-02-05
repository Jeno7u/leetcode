package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil { //  потому что subRoot это минимум одна нода
        return false
    }
    if isSame(root, subRoot) { //  проверка на одинаковость
        return true 
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) 
}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil{
        return true
    } else if root == nil || subRoot == nil {
        return false
    }

    if root.Val != subRoot.Val {
        return false
    }

    return isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
}