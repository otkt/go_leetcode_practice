// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/
package main

//lint:ignore U1000 Ignore
func minimumArea(grid [][]int) int {
	bigx := len(grid)
	bigy := len(grid[0])
	lx := bigx
	rx := -1
	uy := bigy
	dy := -1
	for y, arr := range grid {
		for x, val := range arr {
			if val == 0 {
				continue
			}
			if x < lx {
				lx = x
			}
			if y < uy {
				uy = y
			}
			if x > rx {
				rx = x
			}
			if y > dy {
				dy = y
			}

		}
	}

	x_dist := max(rx-lx+1, 1)
	y_dist := max(dy-uy+1, 1)
	return y_dist * x_dist
}

// func main() {
// 	matrix := [][]int{
// 		{0, 0, 0},
// 		{0, 1, 1},
// 		{1, 0, 0},
// 		{0, 1, 1},
// 	}
// 	a := minimumArea(matrix)
// 	print(a)
// }
