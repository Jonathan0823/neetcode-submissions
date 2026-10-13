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

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		maxDiameter = max(maxDiameter, leftHeight + rightHeight)

		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)

	return maxDiameter
}
