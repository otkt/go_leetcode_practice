package main

// func main() {
// 	input := []int{3, 2, 4}
// 	target := 6
// 	output := twoSum(input, target)
// 	print(output, target)
// }

//lint:ignore U1000 Ignore unused function temporarily for debugging

func twoSum(nums []int, target int) []int {
	value_to_index := make(map[int]int)

	for i, first := range nums {
		req := target - first
		value_to_index[req] = i
		for j, second := range nums {
			if i == j {
				continue
			}

			if val, ok := value_to_index[second]; ok {
				return []int{val, j}
			} else {
				req := target - second
				value_to_index[req] = j
			}
		}
	}
	return []int{-1, -1}
}
