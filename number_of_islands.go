package main

import (
	"encoding/json"
	"errors"
)

func getgrid() [][]byte {
	// Given input string
	input := `[["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]`

	var arr [][]string
	err := json.Unmarshal([]byte(input), &arr)
	if err != nil {
		panic(err)
	}

	var result [][]byte
	for _, row := range arr {
		var byteRow []byte
		for _, str := range row {
			byteRow = append(byteRow, str[0]) // assuming each string is a single character "0" or "1"
		}
		result = append(result, byteRow)
	}

	return result
}

// func main() {
// 	grid := getgrid()
// 	islands := numIslands(grid)
// 	print(islands)
// }

func numIslands(grid [][]byte) int {
	make_byte_array_non_retarded(&grid)
	g := *NewGrid(grid)
	island_count := g.paint_islands()
	return int(island_count)
}

func make_byte_array_non_retarded(arr *[][]byte) {
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

type CoordinateQueue struct {
	items []Coordinate
}

func (cq *CoordinateQueue) enqueue(g *Grid, x, y int) bool {
	coordinate := Coordinate{x: x, y: y, val: &g.arr[x][y]}
	for _, item := range cq.items {
		if item.x == coordinate.x && item.y == coordinate.y {
			return false
		}
	}
	cq.items = append(cq.items, coordinate)
	return true
}

func (cq *CoordinateQueue) dequeue() (Coordinate, error) {
	if len(cq.items) < 1 {
		return Coordinate{}, errors.New("queue is empty")
	}
	popped := cq.items[0]
	cq.items = cq.items[1:]
	return popped, nil
}

func (g *Grid) valid_neighboors(x, y int) []Coordinate {
	list := []Coordinate{}
	if x > 0 {
		if g.arr[x-1][y] == 1 {
			list = append(list, Coordinate{x: x - 1, y: y, val: &g.arr[x-1][y]})
		}
	}
	if y > 0 {
		if g.arr[x][y-1] == 1 {
			list = append(list, Coordinate{x: x, y: y - 1, val: &g.arr[x][y-1]})
		}
	}
	if x < g.m-1 {
		if g.arr[x+1][y] == 1 {
			list = append(list, Coordinate{x: x + 1, y: y, val: &g.arr[x+1][y]})
		}
	}
	if y < g.n-1 {
		if g.arr[x][y+1] == 1 {
			list = append(list, Coordinate{x: x, y: y + 1, val: &g.arr[x][y+1]})
		}
	}
	return list
}

func (g *Grid) mark_neighboors(val byte, cq *CoordinateQueue) {
	cont := true

	for cont {
		popped, err := cq.dequeue()
		if err == nil {
			*popped.val = val
			for _, n := range g.valid_neighboors(popped.x, popped.y) {
				if *n.val == 1 {
					cq.enqueue(g, n.x, n.y)
				}
			}
		} else {
			cont = false
		}
	}

}

func (g *Grid) get_coordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y, val: &g.arr[x][y]}
}

func (g *Grid) paint_islands() int32 {
	var marking byte = 1
	var islands int32 = 0
	var coordinate_queue CoordinateQueue = CoordinateQueue{}
	for i := range g.m {
		for j := range g.n {
			if g.arr[i][j] == 1 {
				marking++
				islands++
				if marking == 255 {
					marking = 2
				}
				coord := g.get_coordinate(i, j)
				coordinate_queue.enqueue(g, coord.x, coord.y)
				g.mark_neighboors(marking, &coordinate_queue)
			}
		}
	}
	return islands
}
