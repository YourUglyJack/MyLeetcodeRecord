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


 func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next  
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	// 下面会报错，因为用了 fast=fast.Next.Next 又没有保证 fast.Next != nil
	// 可以直接判断fast, 因为fast能走，那slow一定能走
	// for slow != nil && slow.Next != nil && fast != nil  {
	// 	slow = slow.Next
	// 	fast = fast.Next.Next
	// 	if slow == fast {
	// 		return true
	// 	}
	// }
	return false
 }