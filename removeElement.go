package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	for i := range nums {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func main() {
	nums := []int{3,2,2,3}
	val := 3

	length := removeElement(nums, val)
	fmt.Println(nums, length)
}
