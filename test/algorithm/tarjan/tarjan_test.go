package tarjan

import (
	"fmt"
	"testing"
)

type Vertex int

// case1 explanation : https://frankma.netlify.app/notes/programming/used-algorithms
func TestCase1(t *testing.T) {
	adjacencyMap := map[Vertex][]Vertex{
		0: {1},
		1: {2},
		2: {0},
		3: {4, 7},
		4: {5},
		5: {6, 0},
		6: {0, 2, 4},
		7: {5, 3},
	}

	acct := newSccAcct(0, 7, 0)

	//for v := range adjacencyMap {
	//	// 顶点v还未访问
	//	if acct.dfn[v] == -1 {
	//		tarjan(acct, v, adjacencyMap)
	//	}
	//}

	// for debug: the same search sequence as blog's image
	tarjan(acct, Vertex(0), adjacencyMap)
	tarjan(acct, Vertex(3), adjacencyMap)
	fmt.Println(`-------------------------
Correct Answer just for debug sequence:
SCC: [[2 1 0] [6 5 4] [7 3]]
DFN: map[0:0 1:1 2:2 3:3 4:4 5:5 6:6 7:7]
DFS sequence: [0 1 2 3 4 5 6 7]
Low: map[0:0 1:0 2:0 3:3 4:4 5:4 6:4 7:3]
Low Group: map[0:[1 2 0] 3:[3 7] 4:[4 5 6]]
-------------------------`)

	fmt.Printf("SCC: %v\n", acct.SCC)

	fmt.Printf("DFN: %v\n", acct.dfn)
	dfsSeq := make([]Vertex, 0)
	for i := 0; i <= 7; i++ {
		for k, v := range acct.dfn {
			if v == i {
				dfsSeq = append(dfsSeq, k)
			}
		}
	}
	fmt.Printf("DFS sequence: %v\n", dfsSeq)

	fmt.Printf("Low: %v\n", acct.low)
	lowLinkValueGroup := make(map[int][]Vertex)
	for k, v := range acct.low {
		if _, ok := lowLinkValueGroup[v]; !ok {
			lowLinkValueGroup[v] = make([]Vertex, 0)
		}
		lowLinkValueGroup[v] = append(lowLinkValueGroup[v], k)
	}
	fmt.Printf("Low Group: %v\n", lowLinkValueGroup)
}

// case2 explanation : https://byvoid.com/zhs/blog/scc-tarjan/
func TestCase2(t *testing.T) {
	adjacencyMap := map[Vertex][]Vertex{
		1: {3, 2},
		2: {4},
		3: {5, 4},
		4: {6, 1},
		5: {6},
		6: {},
	}

	acct := newSccAcct(1, 6, 1)

	//for v := range adjacencyMap {
	//	// 顶点v还未访问
	//	if acct.dfn[v] == -1 {
	//		tarjan(acct, v, adjacencyMap)
	//	}
	//}

	// for debug
	tarjan(acct, Vertex(1), adjacencyMap)
	fmt.Println(`-------------------------
Correct Answer just for debug sequence:
SCC: [[6] [5] [2 4 3 1]]
DFN: map[1:1 2:6 3:2 4:5 5:3 6:4]
DFS sequence: [1 3 5 6 4 2]
Low: map[1:1 2:5 3:1 4:1 5:3 6:4]
Low Group: map[1:[1 3 4] 3:[5] 4:[6] 5:[2]]
-------------------------`)

	fmt.Printf("SCC: %v\n", acct.SCC)

	fmt.Printf("DFN: %v\n", acct.dfn)
	dfsSeq := make([]Vertex, 0)
	for i := 1; i <= 6; i++ {
		for k, v := range acct.dfn {
			if v == i {
				dfsSeq = append(dfsSeq, k)
			}
		}
	}
	fmt.Printf("DFS sequence: %v\n", dfsSeq)

	fmt.Printf("Low: %v\n", acct.low)
	lowLinkValueGroup := make(map[int][]Vertex)
	for k, v := range acct.low {
		if _, ok := lowLinkValueGroup[v]; !ok {
			lowLinkValueGroup[v] = make([]Vertex, 0)
		}
		lowLinkValueGroup[v] = append(lowLinkValueGroup[v], k)
	}
	fmt.Printf("Low Group: %v\n", lowLinkValueGroup)

}

func tarjan(acct *sccAcct, v Vertex, graph map[Vertex][]Vertex) int {
	// 初始化 DFN(节点访问到的序号) 与 Low(和dfn相同值)
	acct.dfn[v] = acct.gIndex
	acct.low[v] = acct.gIndex
	minIdx := acct.gIndex

	acct.gIndex++
	acct.stack = append(acct.stack, v)

	for _, target := range graph[v] {
		// target节点还未访问
		if acct.dfn[target] == -1 {
			// 继续搜索，找到target邻居的low link value，找到后更新target的low
			lowLinkIdxFromTargetNeighbor := tarjan(acct, target, graph)
			minIdx = min(acct.dfn[v], lowLinkIdxFromTargetNeighbor)
			//minIdx = min(minIdx, lowLinkIdxFromTargetNeighbor)
			// 跟新target的low link value
			acct.low[target] = lowLinkIdxFromTargetNeighbor
			acct.low[v] = minIdx
		} else if acct.inStack(target) {
			// 取最小的 low link value
			minIdx = min(acct.low[v], acct.dfn[target])
		}
	}

	// v顶点的所有邻居的low link value更新完了
	// 判断v顶点的访问序号 是否等于 v邻居的low link value
	// 如果有环路，通过上面的for循环，会让v邻居的low link value一致
	if acct.dfn[v] == minIdx {
		var scc []Vertex
		for {
			v2 := acct.pop()
			scc = append(scc, v2)
			if v2 == v {
				break
			}
		}
		acct.SCC = append(acct.SCC, scc)
	}

	return minIdx
}

type sccAcct struct {
	dfn    map[Vertex]int
	low    map[Vertex]int
	stack  []Vertex
	gIndex int
	SCC    [][]Vertex
}

func newSccAcct(start, size, gIdxSt int) *sccAcct {
	s := &sccAcct{
		dfn:    make(map[Vertex]int),
		low:    make(map[Vertex]int),
		stack:  make([]Vertex, 0),
		gIndex: gIdxSt,
		SCC:    make([][]Vertex, 0),
	}

	for i := start; i <= size; i++ {
		s.dfn[Vertex(i)] = -1
		s.low[Vertex(i)] = 999
	}

	return s
}

func (s *sccAcct) inStack(v Vertex) bool {
	for _, n := range s.stack {
		if n == v {
			return true
		}
	}
	return false
}

func (s *sccAcct) push(n Vertex) {
	s.stack = append(s.stack, n)
}

func (s *sccAcct) pop() Vertex {
	n := len(s.stack)
	if n == 0 {
		fmt.Println("stack is empty")
		return 0
	}
	vertex := s.stack[n-1]
	s.stack = s.stack[:n-1]
	return vertex
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
