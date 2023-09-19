package main

import (
	"fmt"
)

//func findDuplicate(nums []int) int {
//	var numCounter = make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		num := nums[i]
//		_, exists := numCounter[num]
//		if exists {
//			return num
//		} else {
//			numCounter[num] = num
//		}
//	}
//	return 0
//}

func findDuplicate(nums []int) int {
	var visited = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			return nums[i]
		} else {
			visited[nums[i]] = true
		}
	}
	return 0
}

//func findDuplicate(nums []int) int {
//	var visited = make(map[int]bool)
//	for _, num := range nums {
//		if visited[num] {
//			return num
//		} else {
//			visited[num] = true
//		}
//	}
//	return 0
//}

//func findDuplicate(nums []int) int {
//	var numCounter = make(map[int]int)
//	for _, num := range nums {
//		_, exists := numCounter[num]
//		if exists {
//			return num
//		} else {
//			numCounter[num] = num
//		}
//	}
//	return 0
//}

func main() {
	nums := []int{1, 3, 4, 2, 2}

	duplicateNum := findDuplicate(nums)
	fmt.Println(duplicateNum)
}
