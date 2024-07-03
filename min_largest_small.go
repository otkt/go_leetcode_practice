// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/

package main

import (
	"sort"
)

//lint:ignore U1000 I
func minDifference(nums []int) int {
	l := len(nums)
	// if less then 4 all number can always be equal
	if l < 5 {
		return 0
	}
	// else there are 4 worst cases, choose minimum
	diff := l - 4
	sorted := nums
	sort.Ints(sorted)
	choices := []int{}
	skip := 0
	for i := 0; i < 4; i++ {
		choices = append(choices, -1*sorted[i]+sorted[i+diff])
		skip += diff
	}
	min := choices[0]
	for _, c := range choices {
		if c < min {
			min = c
		}
	}
	return min
}

// func main() {
// 	minDifference([]int{5, 4, 3, 2, 1, 6, 7, 8, 9})
// }
