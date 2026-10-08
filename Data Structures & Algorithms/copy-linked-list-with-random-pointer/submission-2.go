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
	curr := head
	for curr != nil{ 
		nodeMap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head 
	for curr != nil { 
		nodeMap[curr].Next = nodeMap[curr.Next]
		nodeMap[curr].Random = nodeMap[curr.Random]
		curr = curr.Next
	}

	return nodeMap[head]
    
}
