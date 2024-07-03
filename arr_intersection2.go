// https://leetcode.com/problems/intersection-of-two-arrays-ii/

package main

//lint:ignore U1000 I
func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)

	for _, n := range nums1 {
		map1[n]++
	}
	var intersection []int
	for _, n := range nums2 {
		n1Count, ok := map1[n]
		if ok && n1Count > 0 {
			intersection = append(intersection, n)
			map1[n]--
		}
	}
	return intersection
}

// func main() {
// 	arr1 := []int{1, 2, 2, 1}
// 	arr2 := []int{2, 2}
// 	res := intersect(arr1, arr2)
// 	print(res)
// }
