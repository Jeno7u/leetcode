package main

 type Node struct {
	Val int
	Neighbors []*Node
 }

// dfs с запоминанием какие мы уже откланировали
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    mapCloned := map[*Node]*Node{}

    var cloning func(node *Node) *Node
    cloning = func(node *Node) *Node {
        if cloned, ok := mapCloned[node]; ok {
            return cloned
        }

        cloned := &Node{Val: node.Val, Neighbors: []*Node{}}
        mapCloned[node] = cloned
        for _, neighbor := range node.Neighbors {
            cloned.Neighbors = append(cloned.Neighbors, cloning(neighbor))
        }
        return cloned
    }
    cloning(node)
    return mapCloned[node]
}