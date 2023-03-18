package util

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{data: []int{0}}
}

// 每次插入最后面，然后从下往上比较节点，是否需要交换
func (heap *MinHeap) insert(data int) {
	heap.data = append(heap.data, data)
	idx := len(heap.data) - 1
	// 只要当前节点比父节点小，则需要交换它两
	for idx > 1 && heap.data[idx] < heap.data[idx/2] {
		heap.data[idx], heap.data[idx/2] = heap.data[idx/2], heap.data[idx]
		idx /= 2
	}
}

// 把最后一个元素放到min位置后，再往下交换
func (heap *MinHeap) removeMin() int {
	// 获取堆中的最小元素
	min := heap.data[1]
	// 将堆中最后一个元素移动到根节点位置
	last := heap.data[len(heap.data)-1]
	heap.data = heap.data[:len(heap.data)-1]
	if len(heap.data) > 1 {
		heap.data[1] = last
		// 对根节点进行下移操作，以维护最小堆结构
		heap.down(1)
	}
	return min
}

// 将某个节点下移，以维护最小堆结构
func (heap *MinHeap) down(idx int) {
	childIdx := idx * 2
	if childIdx > len(heap.data)-1 {
		return
	}
	if childIdx+1 <= len(heap.data)-1 && heap.data[childIdx+1] < heap.data[childIdx] {
		childIdx++
	}
	if heap.data[idx] > heap.data[childIdx] {
		heap.data[idx], heap.data[childIdx] = heap.data[childIdx], heap.data[idx]
		heap.down(childIdx)
	}
}
