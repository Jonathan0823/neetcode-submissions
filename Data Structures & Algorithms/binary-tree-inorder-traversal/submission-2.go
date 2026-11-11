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

	var dfs func(*TreeNode) 
	dfs = func(node *TreeNode) {
		if node == nil { 
			return
		}

		if node.Left != nil { 
			dfs(node.Left)
		}
		res = append(res, node.Val)
		if node.Right != nil { 
			dfs(node.Right)
		}
	}

	dfs(root)

	return res
}
