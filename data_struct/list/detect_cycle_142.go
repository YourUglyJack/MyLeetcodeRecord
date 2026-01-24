package list

type ListNode struct {
	Val int
	Next *ListNode
 }
// 找到相遇点，然后让slow回到head，slow和fast以速度1行走，相遇点就是入口


 // slow 走了 k，那 fast 一定走了 2k，2k-k=k，k是环长度的整数倍
 // 假设 入口 -> 相遇 = m
 // head 离 环入口 是 k-m
 // 入口到相遇是 k-m，我现在在相遇点，环长是k，走了m到相遇点，那想回到起点就是，k-m
func detectCycle(head *ListNode) *ListNode {
    // 检测环
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {  // 相遇了
			break
		}
	}
	if fast == nil ||  fast.Next == nil {  // 没有环
		return nil
	}

	slow = head
	// for fast != nil && fast.Next != nil {  // 这个判断是错的，如果相遇点刚好就是在入口，但是满足循环判断，f s会同时向前，所以最后返回的是，入口的下一个
	// 	fast = fast.Next
	// 	slow = slow.Next
	// 	if fast == slow {
	// 		return slow
	// 	}
	// }
	// 已经判断有环了，所以肯定能一直走
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}