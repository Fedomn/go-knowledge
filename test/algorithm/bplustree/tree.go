package bplustree

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type bplustree struct {
	root *node
	t    int // minimum degree
}

func newBtree(miniDegree int) *bplustree {
	return &bplustree{
		root: nil,
		t:    miniDegree,
	}
}

func (b *bplustree) insert(key int) {
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

func (b *bplustree) remove(key int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("remove key %d\n", key)
			fmt.Printf("%v\n", err)
			// fmt.Printf("%v\n", b.breadthFirstDraw())
			panic(err)
		}
	}()
	if b.root == nil {
		return
	}

	b.root.remove(key)

	if b.root.count == 0 {
		if b.root.isLeaf {
			b.root = nil
		} else {
			b.root = b.root.children[0]
		}
	}
}

func (b *bplustree) preOrderTraversal() []int {
	result := make([]int, 0)
	cursor := b.root
	b.preOrderWalk(cursor, &result)
	return result
}

func (b *bplustree) preOrderTraversalBy(rootNode *node) []int {
	result := make([]int, 0)
	cursor := rootNode
	b.preOrderWalk(cursor, &result)
	return result
}

func (b *bplustree) preOrderWalk(cursor *node, result *[]int) {
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
// important: m in B-Sketcher is bplustree.t*2
func (b *bplustree) breadthFirstWalk() []*node {
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

func (b *bplustree) breadthFirstDraw() string {
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

func (b *bplustree) checkValidity() {
	walkNodeList := b.breadthFirstWalk()
	for _, node := range walkNodeList {
		if node == nil {
			continue
		}
		for idx := 0; idx < node.count; idx++ {
			key := node.keys[idx]
			leftKeys := b.preOrderTraversalBy(node.children[idx])
			for _, lk := range leftKeys {
				if key > lk {
					continue
				}
				fmt.Printf("Encounter invalid leftChild key: %d for key: %d in parent node: %v\n", lk, key, node.keys)
				fmt.Printf("BFS:\n%s", b.breadthFirstDraw())
				panic("Encounter invalid key")
			}
			if idx != len(node.keys)-1 {
				rightKeys := b.preOrderTraversalBy(node.children[idx+1])
				for _, rk := range rightKeys {
					// changed for bplustree
					if key <= rk {
						continue
					}
					fmt.Printf("Encounter invalid rightChild key: %d for key: %d in parent node: %v\n", rk, key, node.keys)
					fmt.Printf("BFS:\n%s", b.breadthFirstDraw())
					panic("Encounter invalid key")
				}
			}
		}
	}
}

// changed for bplustree
func (b *bplustree) checkNextPointer() {
	walkNodes := b.breadthFirstWalk()
	walkLeafKeys := make([]int, 0)
	for _, n := range walkNodes {
		if n == nil {
			continue
		}
		if n.isLeaf {
			for i := 0; i < n.count; i++ {
				walkLeafKeys = append(walkLeafKeys, n.keys[i])
			}
		}
	}
	key := b.firstKey()
	scanKeys := b.scanKeysFrom(key)

	if !reflect.DeepEqual(walkLeafKeys, scanKeys) {
		fmt.Printf("walk leaf keys: %v\n", walkLeafKeys)
		fmt.Printf("scan from first key: %d\n", key)
		fmt.Printf("scan keys: %v\n", scanKeys)
		fmt.Printf("BFS:\n%s\n", b.breadthFirstDraw())
		panic("Encounter invalid next pointer")
	}
}

// changed for bplustree
func (b *bplustree) scanKeysFrom(key int) []int {
	if b.root == nil {
		return nil
	}
	result := make([]int, 0)
	if node, found := b.root.search(key); found {
		for i := 0; i < node.count; i++ {
			result = append(result, node.keys[i])
		}
		nextNode := node.next
		for nextNode != nil {
			for i := 0; i < nextNode.count; i++ {
				result = append(result, nextNode.keys[i])
			}
			nextNode = nextNode.next
		}
		return result
	} else {
		return result
	}
}

// changed for bplustree
func (b *bplustree) firstKey() int {
	if b.root == nil {
		return -1
	}

	return b.root.firstKey()
}
