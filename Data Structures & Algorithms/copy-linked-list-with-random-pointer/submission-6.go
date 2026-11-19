/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    copyMap := make(map[*Node]*Node)
	cur := head

	for cur != nil { 
		copyMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil { 
		copyMap[cur].Next = copyMap[cur.Next]
		copyMap[cur].Random = copyMap[cur.Random]
		cur = cur.Next
 	}

	return copyMap[head]
}
