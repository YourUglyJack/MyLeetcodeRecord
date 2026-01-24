package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


func visit(node *TreeNode) {
	fmt.Println(node.Val)
}

func preTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	visit(root)
	preTraverse(root.Left)
	preTraverse(root.Right)
}

func midTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	midTraverse(root.Left)
	visit(root)
	midTraverse(root.Right)
}

func build(nums []int, l, r int) *TreeNode{
	if l > r {
		return nil
	}

	mid := (l + r) / 2 // 找中间
	root := &TreeNode{Val: nums[mid]}

	root.Left = build(nums, l, mid-1)
	root.Right = build(nums, mid+1, r)

	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
    return build(nums, 0, len(nums)-1)
}

func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {  // 找到了插入的位置
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertBST(root.Left, val) // 去左边找插入的位置
	} else if val > root.Val {
		root.Right = insertBST(root.Right, val)
	}
	return root
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func main() {
	nums := []int{1,2,3,4,5}
	midTraverse(sortedArrayToBST(nums))
}