package main

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// делаем preorder и сериализуем subtrees для того чтобы мы могли их положить в hash map и проверять на наличие одинаковых
// subtrees
func dfs(root *TreeNode, res *[]*TreeNode, subtrees map[string]int) string {
	if root == nil {
		return "nil"
	}
	s := strconv.Itoa(root.Val) + "," + dfs(root.Left, res, subtrees) + "," + dfs(root.Right, res, subtrees)
	if subtrees[s] == 1 {
		*res = append(*res, root)
	}
	subtrees[s]++
	
	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	subtrees := map[string]int{}
	dfs(root, &res, subtrees)
	return res
}
