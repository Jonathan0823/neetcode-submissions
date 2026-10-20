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
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) { 
		if node == nil { 
			return
		}

		if len(res) == depth { 
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)
	return res
}
