package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct { 
	Val int
	Next *ListNode
 }

 // find Nth from end
 // NOTE: fast 先走n步，当fast走到nil，slow就刚好在倒数n的位置
 func findNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
 }


 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // NOTE: 加dummy后，找倒数第n+1个结点
	dummy := &ListNode{
		Val: -1,
		Next: head,
	}

	targetNode := findNthFromEnd(dummy, n+1)

	targetNode.Next = targetNode.Next.Next

	// dummy.Next 指向原linkedlist的head
	// 直接head是错的，因为head指向head，从来没变过，如果删除head，返回head 并不能得到nil
	return dummy.Next
 }


 func removeNthFromEndV2(head *ListNode, n int) *ListNode {
    // NOTE: 加dummy后，找倒数第n+1个结点
	dummyNode := &ListNode{Next: head}
	fast, slow := dummyNode, dummyNode

	for i := 0; i < n+1; i++ {
		// 先让fast前进n步 
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
 }