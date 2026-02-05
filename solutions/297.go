package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func ConstructorAnother() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var serialization func(root *TreeNode) string 
    serialization = func(root *TreeNode) string {
        if root == nil {
            return "n"
        }
        return strconv.Itoa(root.Val) + "," + serialization(root.Left) + "," + serialization(root.Right)
    }
    return serialization(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    nodeValues := strings.Split(data, ",")
    i := 0

    var construct func() *TreeNode 
    construct = func() *TreeNode {
        nodeVal := nodeValues[i]
        i++
        if nodeVal == "n" {
            return nil
        }
        nodeValConverted, _ := strconv.Atoi(nodeVal)
        node := &TreeNode{Val: nodeValConverted}
        node.Left = construct()
        node.Right = construct()
        return node
    }

    return construct()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */