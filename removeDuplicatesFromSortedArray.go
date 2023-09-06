package main

import "fmt"

// Array slicing
//func removeDuplicates(nums []int) int {
//	for i := 0; i < len(nums)-1; {
//		if nums[i] == nums[i+1] {
//			nums = append(nums[:i+1], nums[i+2:]...)
//		} else {
//			i++
//		}
//	}
//	return len(nums)
//}

// one loop
func removeDuplicates(nums []int) int {
	i := 0
	for j := range nums {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2}
	//nums := []int{0,0,1,1,1,2,2,3,3,4}
	//nums := []int{1, 1}

	length := removeDuplicates(nums)
	println(length)
	fmt.Println(nums)
}
