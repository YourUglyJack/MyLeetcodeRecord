"""
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
！示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
！示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
！示例 3:
输入: s = "pwwkew"
输出: 3
！解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
！提示：
0 <= s.length <= 5 * 10⁴
s 由英文字母、数字、符号和空格组成
！Related Topics 哈希表 字符串 滑动窗口 👍 8228 👎 0
"""


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def lengthOfLongestSubstring(self, s):
        if not s:
            return 0
        n = len(s)
        left = 0
        max_len = 0
        cur_len = 0
        lookup = set()
        for i in range(n):
            cur_len += 1
            while s[i] in lookup:  # 如果当前字符和前一个字符一样 连续删  pwwk
                lookup.remove(s[left])
                left += 1
                cur_len-=1
            lookup.add(s[i])
            max_len = max(max_len,cur_len)
        return max_len
# leetcode submit region end(Prohibit modification and deletion)
s = "abcabcbb"
sol = Solution()
sol.lengthOfLongestSubstring(s)
