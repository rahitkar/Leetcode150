package main

import (
	"fmt"
)

// With out extra variables
//func majorityElement(nums []int) int {
//	sort.Ints(nums)
//	c := 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i-1] == nums[i] {
//			c++
//			if c > len(nums) / 2 {
//				return nums[i]
//			}
//		} else {
//			c = 1
//		}
//	}
//	return nums[0]
//}

// With extra variables but more readable
//func majorityElement(nums []int) int {
//	sort.Ints(nums)
//	c := 1
//	prevNum := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if prevNum == nums[i] {
//			c++
//			if c > len(nums) / 2 {
//				return prevNum
//			}
//		} else {
//			c = 1
//			prevNum = nums[i]
//		}
//	}
//	return prevNum
//}

// With map
func majorityElement(nums []int) int {
	var counter = make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	for c := range counter {
		if counter[c] > len(nums) / 2 {
			return c
		}
	}
	return -1
}

func main() {
	//nums := []int{3, 2, 3}
	nums := []int{1, 1, 2, 2, 3, 1, 2, 2}
	//nums := []int{1}

	majorityEle := majorityElement(nums)
	fmt.Println(majorityEle)
}
