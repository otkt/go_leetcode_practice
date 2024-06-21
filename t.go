package main

func maiasdsn() {
	// grid := [][]string{
	// 	{"1", "1", "1", "1", "0"},
	// 	{"1", "1", "0", "1", "0"},
	// 	{"1", "1", "0", "0", "0"},
	// 	{"0", "0", "0", "0", "0"},
	// }
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	g := *NewGrid(grid)
	island_count := g.paint_islands()
	print(int(island_count))
}

func make_byte_array_non_retarted(arr *[][]byte) {
	m := len(*arr)
	n := len((*arr)[0])
	for i := range m {
		for j := range n {
			if (*arr)[i][j] == 48 {
				(*arr)[i][j] = 0
			} else {
				(*arr)[i][j] = 1
			}
		}
	}

}

type Grid struct {
	arr [][]byte
	m   int
	n   int
}

func NewGrid(arr [][]byte) *Grid {
	m := len(arr)
	n := len(arr[0])
	return &Grid{arr: arr, m: m, n: n}
}

type Coordinate struct {
	x   int
	y   int
	val *byte
}

func (g Grid) valid_neighboors(x, y int) []Coordinate {
	list := []Coordinate{}
	if x > 0 {
		list = append(list, Coordinate{x: x - 1, y: y, val: &g.arr[x-1][y]})
	}
	if y > 0 {
		list = append(list, Coordinate{x: x, y: y - 1, val: &g.arr[x][y-1]})
	}
	if x < g.m-1 {
		list = append(list, Coordinate{x: x + 1, y: y, val: &g.arr[x+1][y]})
	}
	if y < g.n-1 {
		list = append(list, Coordinate{x: x, y: y + 1, val: &g.arr[x][y+1]})
	}
	return list
}

func (g Grid) mark_neighboors(c Coordinate, val byte) {
	*c.val = val
	for _, n := range g.valid_neighboors(c.x, c.y) {
		if *n.val == 1 {
			g.mark_neighboors(n, val)
		}
	}
}

func (g Grid) get_coordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y, val: &g.arr[x][y]}
}

func (g Grid) paint_islands() byte {
	var marking byte = 1
	for i := range g.m {
		for j := range g.n {
			if g.arr[i][j] == 1 {
				marking++
				g.mark_neighboors(g.get_coordinate(i, j), marking)
			}
		}
	}
	return marking - 1
}
