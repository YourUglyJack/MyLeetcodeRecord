package main

import "fmt"

func moveZeroes(nums []int) {
	zeroCount, nonZeroCount := 0, 0 //  当前0的个数
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount] = nums[i]
			nonZeroCount++
		}
	}
	//fmt.Println("quit", zeroCount, nonZeroCount)
	for i := 0; i < zeroCount; i++ {
		//fmt.Println(nonZeroCount - 1 + i)
		nums[nonZeroCount+i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
