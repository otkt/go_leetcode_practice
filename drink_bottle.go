// https://leetcode.com/problems/water-bottles/

package main

//lint:ignore U1000 Ignore
func numWaterBottles(numBottles int, numExchange int) int {
	var drunk int
	var emptyBottles int
	for numBottles > 0 {
		drunk += numBottles
		emptyBottles += numBottles
		numBottles = 0
		// if exchangeable go exchange
		if emptyBottles >= numExchange {
			numBottles += emptyBottles / numExchange
			emptyBottles = emptyBottles % numExchange //remainder emptry bottles
		}
	}
	return drunk
}
