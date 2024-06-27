// https://leetcode.com/problems/find-center-of-star-graph/?envType=daily-question&envId=2024-06-27
package main

//lint:ignore U1000 Ignore
func findCenter(edges [][]int) int {
	n0 := edges[0][0]
	n1 := edges[0][1]
	if n0 == edges[1][0] || n0 == edges[1][1] {
		return n0
	}
	if n1 == edges[1][0] || n1 == edges[1][1] {
		return n1
	}
	return 420
}
