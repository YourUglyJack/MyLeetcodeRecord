"""
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 请你将两个数相加，并以相同形式返回一个表示和的链表。
 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
 示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
 示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

 提示：
 每个链表中的节点数在范围 [1, 100] 内
 0 <= Node.val <= 9
 题目数据保证列表表示的数字不含前导零
 Related Topics 递归 链表 数学 👍 8727 👎 0
"""


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# todo 待优化 可以将空缺位置进行补0  这样会快很多！
class Solution(object):
    def listLen(self, lst):
        l = 0
        while lst != None:
            l += 1
            lst = lst.next
        return l

    def cal_v(self, f, v):
        if f:
            v += 1
            f = False
        if v >= 10:
            if v == 10:
                v = 0
            else:
                v = v % 10
            f = True
        return f, v

    def addTwoNumbers(self, l1, l2):
        res = ListNode()
        r = res
        f = False  # 是否进位
        if self.listLen(l1) > self.listLen(l2):
            long = l1
            short = l2
        else:
            long = l2
            short = l1
        # 先处理第一个结点
        i = 0
        while short != None:
            v = short.val + long.val
            f, v = self.cal_v(f, v)
            if i == 0:
                res.val = v
            else:
                res.next = ListNode(v)
                res = res.next
            short = short.next
            long = long.next
            i += 1
        while long != None:
            v = long.val
            f, v = self.cal_v(f, v)
            res.next = ListNode(v)
            res = res.next
            long = long.next
        if f:
            res.next = ListNode(1)
        return r
# leetcode submit region end(Prohibit modification and deletion)
