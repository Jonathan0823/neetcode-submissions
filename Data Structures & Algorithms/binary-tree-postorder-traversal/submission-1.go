/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) { 
		if node == nil { 
			return
		}

		if node.Left != nil { 
			dfs(node.Left)
		}
		if node.Right != nil { 
			dfs(node.Right)
		}

		res = append(res, node.Val)
	}

	dfs(root)
	return res
}
