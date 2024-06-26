// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

//lint:ignore U1000 Ignore
func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}
	VisitInOrder(root, &arr)
	return arr
}

// func main() {
// 	arr := inorderTraversal(nil)
// 	print(arr)
// }
