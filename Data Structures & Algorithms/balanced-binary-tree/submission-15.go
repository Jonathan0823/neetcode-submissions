/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	return countHeight(root) != -1
}

func countHeight(node *TreeNode) int { 
	if node == nil { 
		return 0
	}

	leftHeight := countHeight(node.Left)
	if leftHeight == -1 { 
		return -1
	}

	rightHeight := countHeight(node.Right)
	if rightHeight == -1 { 
		return -1
	}

	dif := leftHeight - rightHeight
	if dif < 0 { 
		dif *= -1
	}
	
	if dif > 1 { 
		return -1
	}

	return 1 + max(countHeight(node.Left), countHeight(node.Right))
}
