package main

func rotateV1(nums []int, k int) {
	if k == 0 {
		return
	}
	// mod
	l := len(nums)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		pos := (i + k) % l
		res[pos] = nums[i]
	}
	for i := 0; i < l; i++ {
		nums[i] = res[i]
	}
}

func rotateV2(nums []int, k int) { // 环状替换
	if k == 0 {
		return
	}

	count, l := 0, len(nums)
	for startIdx := 0; count < l; startIdx++ {
		preValue, curIdx := nums[startIdx], startIdx
		for {
			nextIdx := (curIdx + k) % l
			tmp := nums[nextIdx]
			nums[nextIdx] = preValue
			preValue = tmp
			curIdx = nextIdx
			count++
			if curIdx == startIdx {
				break
			}
		}
	}

}
