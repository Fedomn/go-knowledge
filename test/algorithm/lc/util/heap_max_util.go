package util

//根据 维基百科 的定义，堆 是一种特别的二叉树，满足以下条件的二叉树，可以称之为 堆：
//- 完全二叉树；
//- 每一个节点的值都必须 大于等于或者小于等于 其孩子节点的值。

//最大堆：堆中每一个节点的值 都大于等于 其孩子节点的值。所以最大堆的特性是 堆顶元素（根节点）是堆中的最大值。
//最小堆：堆中每一个节点的值 都小于等于 其孩子节点的值。所以最小堆的特性是 堆顶元素（根节点）是堆中的最小值。

//最大堆是一种树形数据结构，其中每个父节点的值都大于或等于其子节点的值。用数组实现最大堆时，可以按照完全二叉树的方式存储，并符合以下规则：
//
//• 下标从1开始；
//• 如果一个节点的下标为 i，则它的左子节点的下标为 2i，右子节点的下标为 2i+1；
//• 如果一个节点的下标为 i，则它的父节点的下标为 i/2。
//• 如果一个节点的下标为 i，则它是否为叶子节点: i > n/2

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: []int{0}}
}

// 每次插入最后面，然后从下往上比较节点，是否需要交换
func (heap *MaxHeap) insert(val int) {
	heap.data = append(heap.data, val)
	idx := len(heap.data) - 1
	// 只要当前节点比父节点大，则需要交换它两
	for idx > 1 && heap.data[idx] > heap.data[idx/2] {
		heap.data[idx], heap.data[idx/2] = heap.data[idx/2], heap.data[idx]
		idx /= 2
	}
}

// 把最后一个元素放到max位置后，再往下交换
func (heap *MaxHeap) removeMax() int {
	//1.获取堆中的最大元素，并保存到变量max中。
	max := heap.data[1]
	//2.获取堆中的最后一个元素，并保存到变量last中。
	last := heap.data[len(heap.data)-1]
	//3.将堆的大小减1，并将最后一个元素替换为之前保存的最大元素。
	heap.data = heap.data[:len(heap.data)-1]
	if len(heap.data) > 1 {
		heap.data[1] = last
		//4.调用down函数对新的根节点进行下移操作，以恢复堆的最大堆结构。
		heap.down(1)
	}
	return max
}

// 函数用于删除根节点后重建堆
func (heap *MaxHeap) down(idx int) {
	//1.首先计算当前节点的左子节点下标，即2*idx，如果左子节点的下标超出堆的大小，则当前节点没有子节点，直接退出函数。
	childIdx := 2 * idx
	if childIdx > len(heap.data)-1 {
		return
	}
	//2.如果当前节点存在右子节点，并且右子节点的值大于左子节点的值，则选择右子节点为优先级更高的子节点。
	if childIdx+1 <= len(heap.data)-1 && heap.data[childIdx+1] > heap.data[childIdx] {
		childIdx++
	}
	//3.如果当前节点的值小于优先级更高的子节点的值，则将当前节点的值和该子节点的值交换，并对该子节点进行递归down操作。
	if heap.data[idx] < heap.data[childIdx] {
		heap.data[idx], heap.data[childIdx] = heap.data[childIdx], heap.data[idx]
		// 交换完后，继续往下交换
		heap.down(childIdx)
	}
	//4.否则，当前节点不需要下移，函数退出。
	//通过这些步骤，down函数可以有效地将根节点下移，以重新构造最大堆。
}
