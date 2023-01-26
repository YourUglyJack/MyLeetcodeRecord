package main

// 方法一：找到正负的分界线，利用快排思维解题
// 方法二：双指针,最大值要么在最右边要么在最左边

func sortedSquares(nums []int) []int {
	l := len(nums)
	left, right := 0, l-1
	res := make([]int, l)
	for pos := l - 1; pos >= 0; pos-- {
		ls, rs := nums[left]*nums[left], nums[right]*nums[right]
		if ls >= rs {
			res[pos] = ls
			left++
		} else {
			res[pos] = rs
			right--
		}
	}
	return res
}
