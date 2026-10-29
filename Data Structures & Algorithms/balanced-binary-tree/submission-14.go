/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil { 
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) { 
		return false
	}

	leftHeight := countHeight(root.Left)
	rightHeight := countHeight(root.Right)
	
	dif := leftHeight - rightHeight
	if dif < 0 { 
		dif *= -1
	}

	if dif > 1 { 
		return false
	}
	
	return true
}

func countHeight(node *TreeNode) int { 
	if node == nil { 
		return 0
	}

	return 1 + max(countHeight(node.Left), countHeight(node.Right))
}
