package main

import "fmt"

//Given an array of integers nums which is sorted in ascending order, and an integer target,
//write a function to search target in nums.If target exists, then return its index.
//Otherwise, return -1. You must write an algorithm with O(log n) runtime complexity.

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2 // 存在int相加溢出的情况，可以考虑换一个公式
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1

}

func main() {
	t := []int{-1, 0, 3, 5, 9, 12} // slice
	idx := search(t, 9)
	fmt.Println(idx)
}
