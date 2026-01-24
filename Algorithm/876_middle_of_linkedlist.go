package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLen(head *ListNode) int {
	var p = head
	var lstLen = 0
	for p != nil {
		lstLen += 1
		p = p.Next
	}
	return lstLen
}

func middleNode(head *ListNode) *ListNode {
	var lstLen = getLen(head)
	var mid, curCount = lstLen / 2, 0
	var p = new(ListNode)
	var ret = p
	for head != nil {
		curCount++
		if curCount >= mid {
			if curCount == mid {
				p.Val = head.Val
				p.Next = head.Next
			} else {
				var q = new(ListNode)
				q.Val = head.Val

			}

		} else {
			head = head.Next
		}
	}
	return ret
}
