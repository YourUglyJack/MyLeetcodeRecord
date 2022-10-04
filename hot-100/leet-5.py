"""
给你一个字符串 s，找到 s 中最长的回文子串。
 示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
 示例 2：
输入：s = "cbbd"
输出："bb"
 提示：
 1 <= s.length <= 1000
 s 仅由数字和英文字母组成
 Related Topics 字符串 动态规划 👍 5768 👎 0
"""


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def longestPalindrome(self, s):
        n = len(s)
        dp = [[False for _ in range(n)] for _ in range(n)]
        start = 0
        end = 0
        max_len = 0
        # dp[i][j] 表示 [i,j]是否是回文子串  如果是 则要看[i+1][j-1]是不是回文子串  所以遍历的时候从下到上 从右到左
        # 该二维数组只更新右上三角
        for i in range(n - 1, -1, -1):
            for j in range(i, n):
                if s[i] == s[j]:
                    if j - i <= 1:
                        dp[i][j] = True
                    elif dp[i + 1][j - 1]:
                        dp[i][j] = True
                if dp[i][j] and j - i + 1 > max_len:
                    max_len = j - i + 1
                    start = i
                    end = j
        return s[start:end + 1]


s = "bb"
sol = Solution()
print(sol.longestPalindrome(s))
# leetcode submit region end(Prohibit modification and deletion)
