package main

import "sort"

//with decreasing sort
func hIndex(citations []int) int {
	hIndex := 0
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := 0; i < len(citations); i++ {
		if i+1 <= citations[i] {
			hIndex = i + 1
		} else {
			return hIndex
		}
	}
	return hIndex
}


/*
1 -> 5
2 -> 3
5 -> 1
*/

func main() {
	//citations := []int{3, 0, 6, 1, 5} // 6, 5, 3, 1, 0 // 0, 1, 3, 5, 6
	citations := []int{1, 2, 3, 5, 1} // 5, 3, 2, 1, 1 // 1, 1, 2, 3, 5
	//citations := []int{1, 3, 1} // 3, 1, 1 // 1, 1, 3
	hIndex := hIndex(citations)
	println(hIndex)
}
