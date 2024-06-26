// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/
package main

//lint:ignore U1000 I
func minimumAverage(nums []int) float64 {
	n := len(nums) / 2
	var avgs []float64
	for i := 0; i < n; i++ {
		var smallest int = nums[0]
		s_i := 0
		b_i := 0
		var biggest int = nums[0]
		var avarage float64
		for i, num := range nums {
			if num < smallest {
				smallest = num
				s_i = i
			}
			if num > biggest {
				biggest = num
				b_i = i
			}

		}
		avarage = float64(biggest+smallest) / 2
		avgs = append(avgs, avarage)
		nums = append(nums[:s_i], nums[s_i+1:]...)
		if b_i > s_i {
			b_i--
		}
		nums = append(nums[:b_i], nums[b_i+1:]...)
	}
	var min_avg float64 = avgs[0]
	for _, avg := range avgs {
		if avg < min_avg {
			min_avg = avg
		}
	}
	return min_avg
}
