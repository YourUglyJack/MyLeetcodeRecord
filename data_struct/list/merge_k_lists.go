package list

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap struct {
	data []*ListNode
}
 
func left(i int) int {
	return 2 * i + 1
}

func right(i int) int {
	return 2 * i +2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (mh *minHeap) Less(i, j int) bool {
	return mh.data[i].Val < mh.data[j].Val
}

func (mh *minHeap) Swap(i, j int) {
	mh.data[i], mh.data[j] = mh.data[j], mh.data[i]
}

func (mh *minHeap) down(i int) {
	n := len(mh.data)
	for {
		smallest := i
		l, r := left(i), right(i)
		
		if l < n && mh.Less(l, smallest) {
			smallest = l
		}
		if r < n && mh.Less(r, smallest) {
			smallest = r
		}

		if smallest == i {
			break
		}

		mh.Swap(i, smallest)
		i = smallest
	}
}

func (mh *minHeap) Pop() *ListNode{
	if len(mh.data) == 0 {
		panic("empty")
	}

	min := mh.data[0]
	mh.data[0] = mh.data[len(mh.data)-1]
	mh.data = mh.data[:len(mh.data)-1]
	if len(mh.data) > 0 {
		mh.down(0)
	}
	return min
}

func (mh *minHeap) Push(x *ListNode) {
	mh.data = append(mh.data, x)
	mh.up(len(mh.data)-1)
}
func (mh *minHeap) up(i int) {
	for i > 0 {  // 不加 i > 0, 当p=0的时候，会死循环
		p := parent(i)
		if !mh.Less(i, p) {
			break
		}
		mh.Swap(p, i)
		i = p
	}
}


func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	heap := &minHeap{
		data: make([]*ListNode, 0),
	}
	for _, head := range lists {
		if head != nil {
			heap.Push(head)
		}
	}
	dummy := &ListNode{}
	p := dummy
	for len(heap.data) > 0 {
		node := heap.Pop()
		p.Next = node
		if node.Next != nil {
			heap.Push(node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}
