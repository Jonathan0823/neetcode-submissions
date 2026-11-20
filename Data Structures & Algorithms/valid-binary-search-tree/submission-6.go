/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max *int) bool
	dfs = func(node *TreeNode, min, max *int) bool { 
		if node == nil { 
			return true
		}

		if min != nil && node.Val <= *min { 
			return false
		}

		if max != nil && node.Val >= *max { 
			return false
		}

		return dfs(node.Left, min, &node.Val) && dfs(node.Right, &node.Val, max)
	}

	return dfs(root, nil, nil)
}


