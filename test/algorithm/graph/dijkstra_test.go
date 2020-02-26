package graph__test

import (
	"fmt"
	"testing"
)

func TestDijkstraMatrix(t *testing.T) {
	maxD := 999
	matrix := [6][6]int{
		{0, 1, 12, maxD, maxD, maxD},
		{maxD, 0, 9, 3, maxD, maxD},
		{maxD, maxD, 0, maxD, 5, maxD},
		{maxD, maxD, 4, maxD, 13, 15},
		{maxD, maxD, maxD, maxD, 0, 4},
		{maxD, maxD, maxD, maxD, maxD, 0},
	}

	// 构造 1号点距离其它点最短路径数组
	dis := [6]int{}
	for i := 0; i < len(dis); i++ {
		dis[i] = matrix[0][i]
	}

	// 构造 记录顶点状态数据结构
	findNode := make([]int, 0)
	leftNode := make([]int, 0)
	findNode = append(findNode, 0)
	for i := 1; i < 6; i++ {
		leftNode = append(leftNode, i)
	}

	for true {
		// 从剩余点中 找出距离1号最近的点 minI
		minV, minI := maxD, 0
		for _, e := range leftNode {
			if dis[e] < minV {
				minV = dis[e]
				minI = e
			}
		}

		// 找出该点的所有 出边 outNode
		outNode := make([]int, 0)
		for i, v := range matrix[minI] {
			if v > 0 && v < maxD {
				outNode = append(outNode, i)
			}
		}

		// 对该点的所有 出边 松弛操作
		for _, e := range outNode {
			if dis[e] > dis[minI]+matrix[minI][e] {
				dis[e] = dis[minI] + matrix[minI][e]
			}
		}

		// 移除找到最小距离的点
		for k, v := range leftNode {
			if v == minI {
				findNode = append(findNode, v)
				leftNode = append(leftNode[:k], leftNode[k+1:]...)
			}
		}

		if len(leftNode) == 0 {
			break
		}
	}

	fmt.Println(dis)      // [0 1 8 4 13 17]
	fmt.Println(findNode) // [0 1 3 2 4 5]
	fmt.Println(leftNode) // []
}

func TestDijkstraAdjacencyList(t *testing.T) {
	maxD := 999
	type Node map[int]int
	adjacencyMap := map[int]Node{
		0: {1: 1, 2: 12},
		1: {2: 9, 3: 3},
		2: {4: 5},
		3: {2: 4, 4: 13, 5: 15},
		4: {5: 4},
	}

	// 构造 1号点距离其它点最短路径数组
	dis := [6]int{0, maxD, maxD, maxD, maxD, maxD}
	for k, v := range adjacencyMap[0] {
		dis[k] = v
	}

	// 构造 记录顶点状态数据结构
	findNode := make([]int, 0)
	leftNode := make([]int, 0)
	findNode = append(findNode, 0)
	for i := 1; i < 6; i++ {
		leftNode = append(leftNode, i)
	}

	for true {
		// 从剩余点中 找出距离1号最近的点 minI
		minV, minI := maxD, 0
		for _, e := range leftNode {
			if dis[e] < minV {
				minV = dis[e]
				minI = e
			}
		}

		// 找出该点的所有 出边 outNode
		outNode := adjacencyMap[minI]

		// 对该点的所有 出边 松弛操作
		for k, v := range outNode {
			if dis[k] > dis[minI]+v {
				dis[k] = dis[minI] + v
			}
		}

		// 移除找到最小距离的点
		for k, v := range leftNode {
			if v == minI {
				findNode = append(findNode, v)
				leftNode = append(leftNode[:k], leftNode[k+1:]...)
			}
		}

		if len(leftNode) == 0 {
			break
		}
	}

	fmt.Println(dis)      // [0 1 8 4 13 17]
	fmt.Println(findNode) // [0 1 3 2 4 5]
	fmt.Println(leftNode) // []
}
