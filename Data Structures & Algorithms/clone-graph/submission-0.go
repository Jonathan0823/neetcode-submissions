/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    nodeMap := make(map[*Node]*Node)

	var dfs func(node*Node) *Node 
	dfs = func(node *Node) *Node { 
		if node == nil { 
			return nil
		}

		if _, found := nodeMap[node]; found { 
			return nodeMap[node]
		}

		copy := &Node{Val: node.Val}
		nodeMap[node] = copy
		for _, nei := range node.Neighbors { 
			copy.Neighbors = append(copy.Neighbors, dfs(nei))
		}
		return copy
	} 

	return dfs(node)
}
