"""
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]
 示例 3：
输入：nums = [1]
输出：[[1]]
 提示：
 1 <= nums.length <= 6
 -10 <= nums[i] <= 10
 nums 中的所有整数 互不相同
 Related Topics 数组 回溯 👍 2245 👎 0

"""


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def permute(self, nums):
        def dfs(nums, size, depth, path, visited, res):
            if depth == size:
                res.append(path[:])
                return
            for i in range(size):
                if not visited[i]:
                    visited[i] = True
                    path.append(nums[i])
                    dfs(nums, size, depth + 1, path, visited, res)
                    visited[i] = False
                    path.pop()

        size = len(nums)
        visited = [False for _ in range(size)]
        res = []
        path = []
        dfs(nums, size, 0, path, visited, res)
        return res


if __name__ == '__main__':
    nums = [1, 2, 3]
    solution = Solution()
    res = solution.permute(nums)
    print(res)
# leetcode submit region end(Prohibit modification and deletion)
