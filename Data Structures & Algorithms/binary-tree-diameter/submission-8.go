/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int { 
		if node == nil { 
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		maxDiameter = max(maxDiameter, left + right)
		return 1 + max(left, right)
	}
	dfs(root)

	return maxDiameter
}

