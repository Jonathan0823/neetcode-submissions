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
	var dfs func(i int, node *TreeNode)
	dfs = func(i int, node *TreeNode) { 
		if node == nil { 
			return 
		}

		if i == len(res) { 
			res = append(res, node.Val)
		}
		dfs(i+1, node.Right)
		dfs(i+1, node.Left)
	}

	dfs(0, root)

	return res
}
