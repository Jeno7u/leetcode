package main

// проверяем, что два root равны. Но суть заключается в том, что так как мы идем симметрично, то мы проверяем root.Left + root.Right и root.Right + root.Left
func isSymmetric(root *TreeNode) bool {
    var validate func(rootL, rootR *TreeNode) bool
    validate = func(rootL, rootR *TreeNode) bool {
        if (rootL == nil && rootR != nil) || (rootL != nil && rootR == nil) {
            return false
        }

        if rootL == nil && rootR == nil {
            return true
        }

        if rootL.Val != rootR.Val {
            return false
        }

        return validate(rootL.Left, rootR.Right) && validate(rootL.Right, rootR.Left)
    }

    return validate(root.Left, root.Right)
}