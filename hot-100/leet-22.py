"""
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 有效字符串需满足：
 左括号必须用相同类型的右括号闭合。
 左括号必须以正确的顺序闭合。
 每个右括号都有一个对应的相同类型的左括号。
 示例 1：
输入：s = "()"
输出：true
 示例 2：
输入：s = "()[]{}"
输出：true
 示例 3：
输入：s = "(]"
输出：false
 提示：
 1 <= s.length <= 10⁴
 s 仅由括号 '()[]{}' 组成

 Related Topics 栈 字符串 👍 3530 👎 0

"""


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def isValid1(self, s):
        stack = list()
        n = len(s)
        res = False
        if n == 1:
            return res
        for i in range(n):
            if s[i] == ')' and stack:
                if stack[-1] == '(':
                    stack.pop()
                else:
                    break
            elif s[i] == ']' and stack:
                if stack[-1] == '[':
                    stack.pop()
                else:
                    break
            elif s[i] == '}' and stack:
                if stack[-1] == '{':
                    stack.pop()
                else:
                    break
            else:
                stack.append(s[i])
        if not stack:
            res = True
        return res

    def isValid(self, s):
        if len(s) % 2 == 1:
            return False
        pairs = {
            ')': '(',
            '}': '{',
            ']': '['
        }
        stack = list()
        for ch in s:
            if ch in pairs:  # 判断ch在不在dic的key里面
                if not stack or stack[-1] != pairs[ch]:
                    return False
                stack.pop()
            else:
                stack.append(ch)
        return not stack


s = "]"
sol = Solution()
sol.isValid(s)
# leetcode submit region end(Prohibit modification and deletion)
