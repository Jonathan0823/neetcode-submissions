/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil { 
		return 0
	}

	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		if node == nil { 
			return 0
		}

		res := 0
		if node.Val >= maxVal { 
			res = 1
		}

		maxVal = max(maxVal, node.Val)
		res += dfs(node.Left, maxVal)
		res += dfs(node.Right, maxVal)

		return res
	}

	return dfs(root, root.Val)


    
}
