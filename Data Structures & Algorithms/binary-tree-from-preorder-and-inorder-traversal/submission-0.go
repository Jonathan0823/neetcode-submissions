/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	indices := make(map[int]int)
	for i, v := range inorder { 
		indices[v] = i
	}

	preIdx := 0
	var dfs func(int, int) *TreeNode
	dfs = func(left, right int) *TreeNode { 
		if left > right { 
			return nil
		}


		rootVal := preorder[preIdx]
		preIdx++

		root := &TreeNode{Val: rootVal}
		mid := indices[rootVal]
		
		root.Left = dfs(left, mid - 1)
		root.Right = dfs(mid + 1, right)

		return root
	}

	return dfs(0, len(preorder) - 1)
}
