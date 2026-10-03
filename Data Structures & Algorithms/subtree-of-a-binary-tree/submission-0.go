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
	
	if isSametree(root, subRoot) { 
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSametree(node *TreeNode, subRoot *TreeNode) bool { 
	if node == nil && subRoot == nil { 
		return true
	}

	if node == nil || subRoot == nil || node.Val != subRoot.Val { 
		return false
	}

	return isSametree(node.Left, subRoot.Left) && isSametree(node.Right, subRoot.Right)
}
