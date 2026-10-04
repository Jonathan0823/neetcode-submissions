/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	cur := root

	for cur != nil {
		if p.Val > cur.Val && q.Val > cur.Val { 
			cur = cur.Right
		} else if p.Val < cur.Val && q.Val < cur.Val { 
			cur = cur.Left
		} else { 
			return cur
		}
	}
	return nil
}
