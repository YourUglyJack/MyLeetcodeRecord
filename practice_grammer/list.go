package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T any](prev *Node[T], elem T, next *Node[T]) *Node[T] {
	return &Node[T]{
		val:  elem,
		prev: prev,
		next: next,
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	curPoint := head
	for i := 1; i < len(arr); i++ {
		curPoint.Next = &ListNode{Val: arr[i]}
		curPoint = curPoint.Next
	}
	return head
}

func printLinkedList(lst *ListNode) {
	if lst == nil {
		fmt.Println("empty list")
		return
	}

	for cur := lst; cur != nil; cur = cur.Next {
		fmt.Printf("%d -> ", cur.Val)
	}
	fmt.Print("nil\n")
}

func findKth(k int, lst *ListNode) (int, bool) {
	if k <= 0 {
		return -1, false
	}

	i := 1
	for cur := lst; cur != nil; cur = cur.Next {
		if i == k {
			return cur.Val, true
		}
		i++
	}
	return -1, false
}

func insertKth(k, val int, head *ListNode) (*ListNode, bool) {
	if k <= 0 || head == nil {
		return head, false
	}

	if k == 1 {
		newHead := &ListNode{
			Val:  val,
			Next: head,
		}
		return newHead, true
	}

	i := 1
	for cur := head; cur != nil; cur = cur.Next {
		// find k-1
		if i == k-1 {
			newNode := &ListNode{
				Val:  val,
				Next: cur.Next,
			}
			cur.Next = newNode
			return head, true
		}
		i++
	}
	return head, false
}

func deleteKth(k int, head *ListNode) (*ListNode, bool) {
	if k <= 0 || head == nil {
		return head, false
	}

	if k == 1 {
		return head.Next, true
	}

	pos := 1
	for cur := head; cur != nil; cur = cur.Next {
		// find k-1
		if pos == k-1 {
			if cur.Next == nil {
				return head, false	//	k 越界
			} else {
				cur.Next = cur.Next.Next
				return head, true
			}
		}
		pos++
	}
	return head, false
}


// TODO: double linkedlist

func main() {
	arr := []int{1,2,3,4,5,6}
	lst := createLinkedList(arr)
	printLinkedList(lst)
	k := 2
	insertVal := -100
	val, ok := findKth(k, lst)
	fmt.Printf("find %d th, success: %v, val: %v\n", k, ok, val)
	_, ok = insertKth(2, insertVal, lst)
	fmt.Printf("insert %d to %d th, success: %v, new list: \n", insertVal, k, ok)
	printLinkedList(lst)
	_, ok = deleteKth(2, lst)
	fmt.Printf("delete %d th, success: %v, new list:\n", k, ok)
	printLinkedList(lst)

}
