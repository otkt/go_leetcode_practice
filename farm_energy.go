// https://leetcode.com/contest/biweekly-contest-134/problems/maximum-points-after-enemy-battles/

package main

import (
	"sort"
)

//lint:ignore U1000 I
func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	var points int64
	sort.Slice(enemyEnergies, func(i, j int) bool {
		return enemyEnergies[i] > enemyEnergies[j]
	})
	last := len(enemyEnergies) - 1

	for i := 0; i < len(enemyEnergies)-1; i++ {

		if currentEnergy >= enemyEnergies[last] {
			points += int64(currentEnergy) / int64(enemyEnergies[last])
			currentEnergy = currentEnergy % enemyEnergies[last]
		}

		if points == 0 {
			return 0
		}
		currentEnergy += enemyEnergies[i]
	}

	if currentEnergy >= enemyEnergies[last] {
		points += int64(currentEnergy) / int64(enemyEnergies[last])
	}

	return points
}

// func main() {
// 	print(maximumPoints([]int{3, 2, 2}, 2))
// }
