package main

import "slices"

// bfs, но после идем по уровням. И также разворачиваем список, если уровень нечетный
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) != 0 {
        level := []int{}
        levelLenght := len(queue)
        for i := 0; i < levelLenght; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if len(result) % 2 != 0 {
            slices.Reverse(level)
        }
        result = append(result, level)
    }

    return result
}