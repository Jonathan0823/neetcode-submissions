/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	cur := head
	for cur != nil { 
		nodeMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil { 
		nodeMap[cur].Next = nodeMap[cur.Next]
		nodeMap[cur].Random = nodeMap[cur.Random]
		cur = cur.Next
	}

	return nodeMap[head]
}
