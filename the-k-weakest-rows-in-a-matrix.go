package main

import (
	"fmt"
	"sort"
)

type row struct {
	idx          int
	soldierCount int
}

func kWeakestRows(mat [][]int, k int) []int {
	soldierCounts := countSoldiers(mat)

	sort.Slice(soldierCounts, getWeak(soldierCounts))

	weakestRows := []int{}
	for i, count := range soldierCounts {
		weakestRows = append(weakestRows, count.idx)
		if i >= k-1 {
			return weakestRows
		}
	}
	return weakestRows
}

func getWeak(soldierCounts []row) func(i int, j int) bool {
	return func(i, j int) bool {
		if soldierCounts[i].soldierCount == soldierCounts[j].soldierCount {
			return soldierCounts[i].idx < soldierCounts[j].idx
		}
		return soldierCounts[i].soldierCount < soldierCounts[j].soldierCount
	}
}

func countSoldiers(mat [][]int) []row {
	soldierCounts := []row{}

	n := len(mat[0])
	for i := 0; i < len(mat); i++ {
		soldiers := 0
		for j := 0; j < n; j++ {
			if mat[i][j] > 0 {
				soldiers++
			} else {
				break
			}
		}
		soldierCounts = append(soldierCounts, row{
			idx:          i,
			soldierCount: soldiers,
		})
	}
	return soldierCounts
}

func main() {
	//mat := [][]int{
	//	{1, 1, 0, 0, 0},
	//	{1, 1, 1, 1, 0},
	//	{1, 0, 0, 0, 0},
	//	{1, 1, 0, 0, 0},
	//	{1, 1, 1, 1, 1},
	//}
	//k := 3

	mat := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
	}
	k := 6

	majorityEle := kWeakestRows(mat, k)
	fmt.Println(majorityEle)
}
