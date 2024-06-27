// https://leetcode.com/problems/path-sum/

package main

//lint:ignore U1000 Ignore
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var isLeaf bool
	if root.Left == nil && root.Right == nil {
		isLeaf = true
	}

	targetSum -= root.Val
	if isLeaf {
		if targetSum == 0 {
			return true
		} else {
			return false
		}
	}
	// not leaf
	if hasPathSum(root.Left, targetSum) {
		return true
	}
	if hasPathSum(root.Right, targetSum) {
		return true
	}
	return false
}
