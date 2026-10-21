/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) { 
		if node == nil { 
			return
		}

		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
		return
	}

	dfs(root)
	return res
}
