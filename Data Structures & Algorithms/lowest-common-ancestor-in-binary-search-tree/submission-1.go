/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val { 
		return lowestCommonAncestor(root.Left, p, q)
	}
	
	if p.Val > root.Val && q.Val > root.Val { 
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
