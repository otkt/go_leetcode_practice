// TODO: Implement parallel merge sort then benchmark
package main

// MergeSort is a recursive function that implements the merge sort algorithm to sort a slice of integers.
// It divides the input slice into two halves, recursively sorts each half, and then merges the two sorted halves
// into a single sorted slice.
//
// Parameters:
//   - arr: A slice of integers to be sorted.
//
// Returns:
//   - A new sorted slice of integers.
//
// Example:
//
//	input := []int{4, 2, 7, 1, 9}
//	output := MergeSort(input)
//	fmt.Println(output) // Output: [1, 2, 4, 7, 9]
func MergeSort(arr []int) []int {
	array_len := len(arr)
	if array_len <= 1 {
		return arr
	}

	divide_index := array_len / 2
	b_sorted := MergeSort(arr[divide_index:])
	a_sorted := MergeSort(arr[:divide_index])
	//Merge
	merged_arr := make([]int, array_len)

	var a_index, b_index, k int

	for a_index < len(a_sorted) && b_index < len(b_sorted) {
		if a_sorted[a_index] <= b_sorted[b_index] {
			merged_arr[k] = a_sorted[a_index]
			a_index++
		} else {
			merged_arr[k] = b_sorted[b_index]
			b_index++
		}
		k++
	}

	for b_index < len(b_sorted) {
		merged_arr[k] = b_sorted[b_index]
		k++
		b_index++
	}

	for a_index < len(a_sorted) {
		merged_arr[k] = a_sorted[a_index]
		k++
		a_index++
	}

	return merged_arr
}
