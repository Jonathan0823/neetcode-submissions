/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { 
		return res
	}

	q := []*TreeNode{root}


	for len(q) > 0{ 
		qLen := len(q)
		level := []int{}

		for i := 0; i < qLen; i++ { 
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)

			if node.Left != nil { 
				q = append(q, node.Left)
			}
			if node.Right != nil { 
				q = append(q, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
    
}
