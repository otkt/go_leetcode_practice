// https://leetcode.com/problems/maximum-total-importance-of-roads/
package main

// no sorting
//
//lint:ignore U1000 Ignore
func maximumImportance(n int, roads [][]int) int64 {
	nodeDegreeMap := make(map[int]int)
	var maxDegree int
	for _, pair := range roads {
		for _, node := range pair {
			nodeDegreeMap[node]++
			if nodeDegreeMap[node] > maxDegree {
				maxDegree = nodeDegreeMap[node]
			}
		}
	}
	pools := make(map[int][]int, maxDegree)
	for node, degree := range nodeDegreeMap {
		pools[degree] = append(pools[degree], node)
	}
	score := n
	nodeScoreMap := make(map[int]int)
	for p := maxDegree; p > 0; p-- {
		nodes, ok := pools[p]
		if ok {
			for _, node := range nodes {
				nodeScoreMap[node] = score
				score--
			}
		}
	}
	var totalImportance int64
	for _, pair := range roads {
		totalImportance += int64(nodeScoreMap[pair[0]]) + int64(nodeScoreMap[pair[1]])
	}
	return totalImportance
}

// func main() {
// 	inp := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}, {4, 0}}
// 	r := maximumImportance(5, inp)
// 	println(r)
// }
