// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/

package main

//lint:ignore U1000 Ignore
func nodesBetweenCriticalPoints(head *ListNode) []int {
	var prevChange int = 0 // -1 decrease  , 0 no change , 1 increase
	var prevVal int = head.Val
	var prevIndex int = 0
	var minDist int = -1
	var firstFoundIndex int = -1
	var lastFoundIndex int = -2
	var currentNode *ListNode = head.Next
	for {
		//check peaks and valleys
		if prevChange == -1 {
			if currentNode.Val > prevVal {
				updateIndexes(&prevIndex, &minDist, &lastFoundIndex, &firstFoundIndex)
			}
		} else if prevChange == 1 {
			if currentNode.Val < prevVal {
				updateIndexes(&prevIndex, &minDist, &lastFoundIndex, &firstFoundIndex)
			}
		}
		// updates for next iter
		if currentNode.Val < prevVal {
			prevChange = -1
		} else if currentNode.Val > prevVal {
			prevChange = 1
		} else {
			prevChange = 0
		}
		prevVal = currentNode.Val
		prevIndex++
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next
	}
	//return
	if lastFoundIndex == -2 || (lastFoundIndex == firstFoundIndex) {
		return []int{-1, -1}
	}
	return []int{minDist, lastFoundIndex - firstFoundIndex}
}

func updateIndexes(prevIndex, minDist, lastIndex, firstIndex *int) {
	if *firstIndex == -1 {
		*firstIndex = *prevIndex
		*lastIndex = *prevIndex
	} else {
		if *minDist == -1 || (*minDist > (*prevIndex - *lastIndex)) {
			*minDist = *prevIndex - *lastIndex
		}
		*lastIndex = *prevIndex
	}
}

// func main() {
// 	head := GetListNode([]int{5, 3, 1, 2, 5, 1, 2})
// 	lis := nodesBetweenCriticalPoints(head)
// 	print(lis)
// }
