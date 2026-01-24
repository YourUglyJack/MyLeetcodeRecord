package list

 type ListNode struct {
	Val int
	Next *ListNode
 }

 // 如果A，B长度一样，那么以相同速率行走，相等的时候，就找到了交点
 // 思路就是，有什么方法能补成一样长
 // A走完去走B，B走完去走A，那长度就一样了
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1 
 }