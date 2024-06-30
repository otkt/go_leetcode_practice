// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/

package main

func findRep(node int, m *map[int]int, path *[]int) int {
	nodeRep := (*m)[node]
	if node == nodeRep {
		return node
	} else {
		*path = append(*path, node)
		return findRep(nodeRep, m, path)
	}
}

func mergeRep(n1, n2 int, m *map[int]int) bool {
	n1_path := []int{}
	n1Rep := findRep(n1, m, &n1_path)
	n2_path := []int{}
	n2Rep := findRep(n2, m, &n2_path)
	if n1Rep != n2Rep {
		(*m)[n2Rep] = n1Rep
		for _, n1er := range n1_path {
			(*m)[n1er] = n1Rep
		}
		for _, n2er := range n2_path {
			(*m)[n2er] = n1Rep
		}
		return true
	}
	return false
}

//lint:ignore U1000 I
func maxNumEdgesToRemove(n int, edges [][]int) int {
	commonSet := make(map[int]int)
	commonEdgeCount := 0
	aEdgeCount := 0
	bEdgeCount := 0
	totalAEdges := 0
	totalBEdges := 0
	totalCEdges := 0
	aIndexes := []int{}
	bIndexes := []int{}
	for el := range n {
		commonSet[el+1] = el + 1
	}
	// C construct disjoint sets from common edges and save them to commonSet
	//startTime := time.Now()
	for i, edge := range edges {
		t := edge[0] // 0:a 1:b 2:c
		if t == 1 {
			aIndexes = append(aIndexes, i)
			totalAEdges++
		} else if t == 2 {
			bIndexes = append(bIndexes, i)
			totalBEdges++
		} else {
			n1 := edge[1]
			n2 := edge[2]
			totalCEdges++
			if mergeRep(n1, n2, &commonSet) {
				commonEdgeCount++
			}
		}
	}
	cRemoved := totalCEdges - commonEdgeCount
	// endTime := time.Now()
	// duration := endTime.Sub(startTime)
	// fmt.Println("Elapsed Time:", duration)
	// A
	// startTime = time.Now()
	aSet := make(map[int]int)
	for el := range n {
		aSet[el+1] = commonSet[el+1]
	}
	for _, index := range aIndexes {
		n1 := edges[index][1]
		n2 := edges[index][2]
		if mergeRep(n1, n2, &aSet) {
			aEdgeCount++
		}
	}
	if aEdgeCount+commonEdgeCount != n-1 {
		return -1
	}
	aRemoved := totalAEdges - aEdgeCount
	// endTime = time.Now()
	// duration = endTime.Sub(startTime)
	// fmt.Println("Elapsed Time:", duration)
	//B
	bSet := make(map[int]int)
	// startTime = time.Now()
	for el := range n {
		bSet[el+1] = commonSet[el+1]
	}
	for _, index := range bIndexes {
		n1 := edges[index][1]
		n2 := edges[index][2]
		if mergeRep(n1, n2, &bSet) {
			bEdgeCount++
		}
	}
	if bEdgeCount+commonEdgeCount != n-1 {
		return -1
	}
	bRemoved := totalBEdges - bEdgeCount
	// endTime = time.Now()
	// duration = endTime.Sub(startTime)
	// fmt.Println("Elapsed Time:", duration)
	return aRemoved + bRemoved + cRemoved
}
