// https://leetcode.com/problems/merge-nodes-in-between-zeros/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//lint:ignore U1000 Ignore
func mergeNodes(head *ListNode) *ListNode {
	modif := &ListNode{}
	merge(head.Next, modif)
	return modif
}

func merge(trav *ListNode, modif *ListNode) {
	if trav == nil {
		return
	}
	modif.Val += trav.Val
	if trav.Val == 0 {
		if trav.Next == nil {
			modif.Next = nil
			return
		}
		newModif := &ListNode{}
		modif.Next = newModif
		merge(trav.Next, modif.Next)
	} else {
		merge(trav.Next, modif)
	}
}

// func main() {
// 	inp := []int{0, 3, 1, 0, 4, 5, 2, 0}
// 	head := GetListNode(inp)
// 	new := mergeNodes(head)
// 	print(new)
// }

func GetListNode(inp []int) *ListNode {
	head := &ListNode{}
	head.LinkSlice(inp)
	return head
}

func (l *ListNode) LinkSlice(arr []int) {
	if len(arr) == 0 {
		return
	}
	l.Val = arr[0]
	arr = arr[1:]
	if len(arr) == 0 {
		return
	}
	l.Next = &ListNode{}
	l.Next.LinkSlice(arr)
}
