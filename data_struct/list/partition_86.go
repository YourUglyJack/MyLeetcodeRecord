package list

type ListNode struct {
	Val int
	Next *ListNode
 }

// 可以借助合并链表
// 构造两个链表，p1都小于x，p2都大于等于x
// 然后合并这两个链表
func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}
	p := head
	p1, p2 := dummy1, dummy2
	// 构造两个新链表，dummy1 都是小于x的，dummy2 都是大于等于x的
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		p = p.Next
	}

	p2.Next = nil // 防止成环
	
	p1.Next = dummy2.Next
	return dummy1.Next
}