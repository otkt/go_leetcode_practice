// https://leetcode.com/problems/three-consecutive-odds/

package main

//lint:ignore U1000 I
func threeConsecutiveOdds(arr []int) bool {
	odd_seq := 0
	for _, val := range arr {
		if val%2 != 0 {
			odd_seq++
		} else {
			odd_seq = 0
		}
		if odd_seq == 3 {
			return true
		}
	}
	return false
}
