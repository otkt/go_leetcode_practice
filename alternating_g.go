//  https://leetcode.com/contest/biweekly-contest-134/problems/alternating-groups-i/

package main

//lint:ignore U1000 I
func numberOfAlternatingGroups1(colors []int) int {
	n := len(colors)
	alternates := 0
	colors = append(colors, colors[0], colors[1])
	for index := range n {
		if colors[index] == 0 {
			if colors[index+1] == 1 && colors[index+2] == 0 {
				alternates++
				continue
			}
		} else {
			if colors[index+1] == 0 && colors[index+2] == 1 {
				alternates++
				continue
			}
		}
	}
	return alternates
}
