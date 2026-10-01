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

		heightLeft := dfs(node.Left)
		heightRight := dfs(node.Right)

		maxDiameter = max(maxDiameter, heightLeft + heightRight)

		return 1 + max(heightLeft, heightRight)
	}

	dfs(root)
	return maxDiameter
}
