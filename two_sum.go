package main

func main() {
	input := []int{2, 7, 11, 15}
	target := 18
	output := twoSum(input, target)
	print(output, target)
}
func twoSum(nums []int, target int) []int {
	for i, first := range nums {
		for j, second := range nums {
			if i == j {
				continue
			}
			if first+second == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
