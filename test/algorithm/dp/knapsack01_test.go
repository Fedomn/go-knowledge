package dp

import (
	"fmt"
	"testing"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type Goods struct {
	Value  int
	Weight int
}

func TestKnapsack01(t *testing.T) {
	// 因为当前决策依赖于上一次求解 所以构造items[0]为{0，0} 才能从items[1]开始
	items := []Goods{
		{0, 0},
		{1, 1},
		{6, 2},
		{18, 5},
		{22, 6},
		{28, 7},
	}
	// 一共有多少个物品
	itemsCount := len(items) - 1
	// 包总容量
	capacity := 11

	// 构造状态矩阵dp[itemsCount+1][capacity+1]
	// 因为数组下标从0开始的原因 要和场景对应上 矩阵的长宽都要+1
	dp := make([][]int, itemsCount+1)
	for i := 0; i < itemsCount+1; i++ {
		dp[i] = make([]int, capacity+1)
	}

	// 方案1：一行一行来构造矩阵
	for i := 1; i <= itemsCount; i++ {
		for c := 1; c <= capacity; c++ {
			w := items[i].Weight
			if c < w {
				dp[i][c] = dp[i-1][c]
			} else {
				v := items[i].Value
				dp[i][c] = max(dp[i-1][c], dp[i-1][c-w]+v)
			}
		}
	}
	fmt.Println(dp[itemsCount][capacity]) // 40

	// 方案2：一列一列来构造矩阵
	for c := 1; c <= capacity; c++ {
		for i := 1; i <= itemsCount; i++ {
			w := items[i].Weight
			if c < w {
				dp[i][c] = dp[i-1][c]
			} else {
				v := items[i].Value
				dp[i][c] = max(dp[i-1][c], dp[i-1][c-w]+v)
			}
		}
	}
	fmt.Println(dp[itemsCount][capacity]) // 40

	// 求最大价值时选择的方案 即找的是 元素的组合
	i := itemsCount
	c := capacity
	for i > 0 && c > 0 {
		if dp[i][c] != dp[i-1][c] {
			fmt.Printf("第%d个物品，空间：%d，价值：%d\n", i, items[i].Weight, items[i].Value)
			c -= items[i].Weight
		}
		i--
	}
	// 第4个物品，空间：6，价值：22
	// 第3个物品，空间：5，价值：18
}
