package practicegrammer


type Heap struct {
	data []int
}

// 取min/max，pop root
// pop后还要维护树的结构
// 取最后一个元素放到root，再进行下沉
// 取最后一个元素是为了，要维护完全二叉树的结构
// down是因为，要满足堆的定义，父节点比孩子 大 / 小

// 插入，插最后一个位置，然后 up
// 插最后是为了满足完全二叉树的结构
// up 是为了，满足堆的定义

func left(i int) int {
	return 2 * i + 1
}

func right(i int) int {
	return 2 * i +2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) up(i int) {
	// 一路和父节点比较，小就往上swap
	for i > 0 {
		p := parent(i)  // 取i的父节点下标
		if h.data[p] <= h.data[i] {
			break
		}
		// 该节点比父节点小，需要up
		// 交换两个节点的值就可以，然后判断新位置是否满足条件
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

func (h *Heap) Push(x int) {
	h.data = append(h.data, x)
	h.up(len(h.data)-1)
}

func (h *Heap) down(i int) {
	n := len(h.data)
	for {
		smallest := i
		l, r := left(i), right(i)

		// 取左孩子之前要判断到底有没有左孩子，所以才 l < n
		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}

		// 取右孩子之前要判断到底有没有右孩子, 所以才 r < n
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}

		// 如果父节点i是最小的，那直接break
		if smallest == i {
			break
		}
		// 否则，交换，然后继续下沉
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]  // 右边先整体求值，再一次性赋给左边
		i = smallest  // smallest 是 l or r，然后继续循环判断这个i是不是满足smallest
	}
}



func (h *Heap) Pop() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	min := h.data[0]
	originLen := len(h.data)
	h.data[0] = h.data[originLen-1]
	h.data = h.data[:originLen-1]

	if len(h.data) > 0 {
		h.down(0)
	}
	return min
}