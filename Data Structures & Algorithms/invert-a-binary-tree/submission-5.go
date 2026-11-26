/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode { 
		if node == nil { 
			return nil
		}

		node.Left, node.Right = node.Right, node.Left
		dfs(node.Left)
		dfs(node.Right)
		return node
	}

	return dfs(root)
}
