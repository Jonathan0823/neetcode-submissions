/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	var res []int
	var dfs func(lvl int, node *TreeNode)
	dfs = func(lvl int, node *TreeNode) { 
		if node == nil { 
			return 
		}

		if lvl == len(res) { 
			res = append(res, node.Val)
		}
		dfs(lvl+1, node.Right)
		dfs(lvl+1, node.Left)
	}

	dfs(0, root)

	return res
}
