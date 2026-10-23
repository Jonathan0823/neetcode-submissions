/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	res := 0

	var dfs func(maxVal int, cur *TreeNode) 
	dfs = func(maxVal int, cur *TreeNode) { 
		if cur == nil { 
			return
		}

		if cur.Val >= maxVal {
			res++
			maxVal = cur.Val
		}

		dfs(maxVal, cur.Left)
		dfs(maxVal, cur.Right)
	}

	dfs(root.Val, root)
	return res
}
