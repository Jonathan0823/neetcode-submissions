/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil { 
		return false
	}
	if isSameTree(root, subRoot) { 
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
    
}

func isSameTree(node *TreeNode, node2 *TreeNode) bool { 
	if node == nil && node2 == nil { 
		return true
	}

	if node == nil || node2 == nil || node.Val != node2.Val { 
		return false
	}

	return isSameTree(node.Left, node2.Left) && isSameTree(node.Right, node2.Right)
}
