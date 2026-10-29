/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	preIdx, inIdx := 0,0

	var dfs func(limit int) *TreeNode
	dfs = func(limit int) *TreeNode { 
		if preIdx >= len(preorder) { 
			return nil
		}

		if inorder[inIdx] == limit { 
			inIdx++
			return nil
		}

		root := &TreeNode{Val: preorder[preIdx]}
		preIdx++
		root.Left = dfs(root.Val)
		root.Right = dfs(limit)

		return root
	}

	return dfs(math.MaxInt)
}
