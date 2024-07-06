// https://leetcode.com/contest/biweekly-contest-134/problems/alternating-groups-ii/
package main

//lint:ignore U1000 I
func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	alternates := 0
	to_append := colors[0 : k-1]
	colors = append(colors, to_append...)
	last_one_was_seq := false
	tail_of_last := -1
	for index := range n {
		if !last_one_was_seq {
			yesseq, tail := checkAlternating(colors[index : index+k])
			if yesseq {
				tail_of_last = tail
				last_one_was_seq = true
				alternates++
			}
		} else { // last one was seq check tail
			if colors[index+k-1] == tail_of_last {
				last_one_was_seq = false
			} else {
				tail_of_last = colors[index+k-1]
				alternates++
			}
		}
	}
	return alternates
}

func checkAlternating(seq []int) (bool, int) {
	last := -1
	for i, el := range seq {
		if i == 0 {
			last = el
			continue
		}
		if el == last {
			return false, -1
		}
		last = el
	}
	return true, seq[len(seq)-1]
}

// func main() {
// 	print(numberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3))

// }
