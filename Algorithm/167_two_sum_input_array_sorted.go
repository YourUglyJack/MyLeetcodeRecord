package main

func twoSum(numbers []int, target int) []int {
	hs := make(map[int]int) // k:value v:index
	for i := 0; i < len(numbers); i++ {
		curValue, needValue := numbers[i], target-numbers[i]
		if idx, ok := hs[needValue]; ok {
			if idx > i+1 {
				return []int{i + 1, idx}
			} else {
				return []int{idx, i + 1}
			}
		} else {
			hs[curValue] = i + 1
		}
	}
	return []int{}
}
