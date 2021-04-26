package v2

import (
	"strconv"
	"strings"
)

type btree struct {
	root *node
	t    int // minimum degree
}

func newBtree(miniDegree int) *btree {
	return &btree{
		root: nil,
		t:    miniDegree,
	}
}

func (b *btree) insert(key int) {
	if b.root == nil {
		b.root = newNode(b.t, true)
		b.root.keys[0] = key
		b.root.count = 1
		return
	}

	// node's keys already full, then grows tree's height
	if b.root.isFull() {
		// ignore duplicated key
		if b.root.hasDuplicatedKey(key) {
			return
		}

		// create new root
		n := newNode(b.t, false)
		n.children[0] = b.root

		// split the old root and move spilledKey to new root n
		n.splitChild(0, b.root)

		// insert key to new node children
		if key < n.keys[0] {
			n.children[0].insertNonFull(key)
		} else {
			n.children[1].insertNonFull(key)
		}

		b.root = n
	} else {
		b.root.insertNonFull(key)
	}
}

func (b *btree) preOrderTraversal() []int {
	result := make([]int, 0)
	cursor := b.root
	b.preOrderWalk(cursor, &result)
	return result
}

func (b *btree) preOrderWalk(cursor *node, result *[]int) {
	if cursor == nil {
		return
	}
	if cursor.isLeaf {
		for i := 0; i < cursor.count; i++ {
			*result = append(*result, cursor.keys[i])
		}
	} else {
		for i := 0; i < cursor.count; i++ {
			*result = append(*result, cursor.keys[i])
			b.preOrderWalk(cursor.children[i], result)
			if i == cursor.count-1 {
				b.preOrderWalk(cursor.children[i+1], result)
			}
		}
	}
}

// for draw
// use tools: https://projects.calebevans.me/b-sketcher/
// important: m in B-Sketcher is btree.t*2
func (b *btree) breadthFirstWalk() []*node {
	bfsResult := make([]*node, 0)
	if b.root == nil {
		return bfsResult
	}
	stack := make([]*node, 0)
	nextLayerStack := make([]*node, 0)
	stack = append(stack, b.root)
	for len(stack) > 0 {
		cursor := stack[0]
		// into bfsResult
		bfsResult = append(bfsResult, cursor)
		if len(stack) >= 1 {
			stack = stack[1:]
		} else {
			stack = []*node{}
		}

		if cursor.isLeaf {
			// impossible to happen
			continue
		} else {
			for i := 0; i < cursor.count; i++ {
				nextLayerStack = append(nextLayerStack, cursor.children[i])
				if i == cursor.count-1 {
					nextLayerStack = append(nextLayerStack, cursor.children[i+1])
				}
			}
		}

		// stack为空，代表这一层出栈完了，可以添加分隔符了
		if len(stack) == 0 {
			bfsResult = append(bfsResult, nil)
			stack = nextLayerStack
			nextLayerStack = make([]*node, 0)
		}
	}
	return bfsResult
}

func (b *btree) breadthFirstDraw() string {
	var result = ""
	walkNodes := b.breadthFirstWalk()
	for _, cursor := range walkNodes {
		if cursor == nil {
			result += "\n"
			continue
		} else {
			for i := 0; i < cursor.count; i++ {
				result += strconv.Itoa(cursor.keys[i])
				if i < cursor.count-1 {
					result += ","
				}
			}
		}
		result += "/"
	}

	// clean unnecessary separator /
	newResult := ""
	splits := strings.Split(result, "\n")
	for _, each := range splits {
		newResult += strings.Trim(each, "/")
		newResult += "\n"
	}
	return strings.TrimSuffix(newResult, "\n")
}
