package main

// когда мы добавляем элементы bst в inorder у нас treeValues
// будут отсортированы. И тогда просто берем k - 1 элемент.
// можно улучшить если перестать идти в глубину если нашли уже значение
func kthSmallest(root *TreeNode, k int) int {
    treeValues := []int{}

    var recc func(root *TreeNode) 
    recc = func(root *TreeNode) {
        if root == nil {
            return
        }
        recc(root.Left)
        treeValues = append(treeValues, root.Val)
        recc(root.Right)
    }
    recc(root)

    return treeValues[k - 1]
}
