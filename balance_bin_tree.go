// https://leetcode.com/problems/balance-a-binary-search-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(arr []int) *TreeNode {
	l := len(arr)
	if l == 1 {
		return &TreeNode{Val: arr[0]}
	}
	mid := l / 2
	mid_val := arr[mid]
	left_nodes := arr[:mid]
	var left_tree *TreeNode
	if len(left_nodes) > 0 {
		left_tree = NewTree(left_nodes)
	}

	right_nodes := arr[mid+1:]
	var righ_tree *TreeNode
	if len(right_nodes) > 0 {
		righ_tree = NewTree(right_nodes)
	}

	return &TreeNode{Val: mid_val, Left: left_tree, Right: righ_tree}
}

func VisitInOrder(root *TreeNode, arr *[]int) {
	if root.Left != nil {
		VisitInOrder(root.Left, arr)
	}
	*arr = append(*arr, root.Val)
	if root.Right != nil {
		VisitInOrder(root.Right, arr)
	}
}

//lint:ignore U1000 Ignore
func balanceBST(root *TreeNode) *TreeNode {
	arr := []int{}
	VisitInOrder(root, &arr)
	return NewTree(arr)
}
