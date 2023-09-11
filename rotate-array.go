package main

import (
	"fmt"
)

// by splitting dosen't work for given input as the K is bigger then len of numns -> nums := []int{1, 2} k := 3
//func rotate(nums []int, k int) {
//	splitPoint := len(nums) - k
//	copy(nums, append(nums[splitPoint:], nums[:splitPoint]...))
//}

// loop + rotate dosen't work for super big slice
//func rotate(nums []int, k int) {
//	tail := len(nums) -1
//	for i := 0; i < k; i++ {
//		copy(nums, append(nums[tail:], nums[:tail]...))
//	}
//}

// brute force loop
func rotate(nums []int, k int)  {
	for i := 0; i < k; i++ {
		temp := nums[len(nums) -1]
		for j := 0; j < len(nums); j++ {
			e := nums[j]
			nums[j] = temp
			temp = e
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	//nums := []int{1, 2}
	//k := 3

	rotate(nums, k)
	fmt.Println(nums)
}
