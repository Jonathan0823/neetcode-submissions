/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil { 
		return res
	}
	queue := []*TreeNode{root}
     
	for len(queue) > 0 { 
		n := len(queue)
		res = append(res, queue[n - 1].Val)
		nextQueue := make([]*TreeNode, 0, n*2)
		for i := 0; i < n; i++ { 
			if queue[i].Left != nil { 
				nextQueue = append(nextQueue, queue[i].Left)
			}

			if queue[i].Right != nil { 
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		queue = nextQueue
	}

	return res
}
