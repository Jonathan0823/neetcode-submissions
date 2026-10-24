/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return valid(root, nil, nil)
}

func valid(node *TreeNode, min, max *int) bool { 
	if node == nil { 
		return true
	}

	if min != nil && node.Val <= *min { 
		return false
	}

	if max != nil && node.Val >= *max { 
		return false
	}

	return valid(node.Left, min, &node.Val) && valid(node.Right, &node.Val, max)

}
