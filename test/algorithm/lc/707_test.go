package lc

// 循环双链表
type MyLinkedList struct {
	dummy *node
}

type node struct {
	val  int
	next *node
	prev *node
}

func Constructor707() MyLinkedList {
	dummy := &node{val: -1}
	dummy.next = dummy
	dummy.prev = dummy
	return MyLinkedList{dummy}
}

func (this *MyLinkedList) Get(index int) int {
	i := 0
	head := this.dummy.next
	for head != this.dummy && i <= index {
		if i == index {
			return head.val
		}
		head = head.next
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	oriHead := this.dummy.next
	n := &node{
		val:  val,
		next: oriHead,
		prev: this.dummy,
	}
	oriHead.prev = n
	this.dummy.next = n
}

func (this *MyLinkedList) AddAtTail(val int) {
	oriTail := this.dummy.prev
	n := &node{
		val:  val,
		next: this.dummy,
		prev: oriTail,
	}
	oriTail.next = n
	this.dummy.prev = n
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	i := 0
	cursor := this.dummy.next
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	for cursor != this.dummy && i <= index {
		if i == index {
			// 原来的往后移一位
			n := &node{
				val:  val,
				next: cursor,
				prev: cursor.prev,
			}
			// 处理原来的前后元素指针
			cursor.prev.next = n
			cursor.prev = n
			return
		}
		cursor = cursor.next
		i++
	}
	if cursor == this.dummy && i == index {
		// 到index遍历完了
		this.AddAtTail(val)
		return
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	i := 0
	cursor := this.dummy.next
	for cursor != this.dummy {
		if i == index {
			cursor.prev.next = cursor.next
			cursor.next.prev = cursor.prev
			return
		}
		cursor = cursor.next
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
