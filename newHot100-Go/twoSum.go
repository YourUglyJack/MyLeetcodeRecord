/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30307
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if idx, ok := seen[diff]; ok{
			return []int{idx, i}
		}
		seen[v] = i
	}
	return nil
}
// @lc code=end



/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

 */

