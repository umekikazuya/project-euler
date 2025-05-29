package main

import (
	"fmt"
	"math"
)

func main() {
	triangle := [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}

	// 1. 三角形の段分のループを回す
	// 2. 各段の要素を上の段との和に置き換える。（隣接する数字が２つある場合は大きい方を評価）
	// 3. 最後の段まで行った際に
	for triangle_i, o := range triangle {
		// 1段目はスキップ。
		if triangle_i == 0 {
			continue
		}
		for item_i, item := range o {
			// 要素を上の段との和に置き換える。（隣接する数字が２つある場合は大きい方を評価）
			// 例外: 段の1番左側の場合は左の隣接数はないためインデックス-1をしない。そのまま返す。
			// 例外: 段の1番右側の場合は上の段ではそのインデックスは存在しないので-1をしない。そのまま返す。
			triangle[triangle_i][item_i] = item + max(triangle, triangle_i, item_i)
		}
	}

	// 数列の段数を取得
	len := len(triangle)
	//　最後の段の一番大きい数を返却
	fmt.Println(maxInSlice(triangle[len-1]))
}

// 要素を上の段との和に置き換える。（隣接する数字が２つある場合は大きい方を評価）
// 例外: 段の1番左側の場合は左の隣接数はないためインデックス-1をしない。そのまま返す。
// 例外: 段の1番右側の場合は上の段ではそのインデックスは存在しないので-1をしない。そのまま返す。
func max(triangle [][]int, triangle_i int, item_i int) int {
	// 例外: 段の1番左側の場合
	if item_i == 0 {
		return triangle[triangle_i-1][item_i]
	}
	// 例外: 段の1番右側の場合
	if len(triangle[triangle_i-1]) <= item_i {
		return triangle[triangle_i-1][item_i-1]
	}
	// 通常時: 要素を上の段との和に置き換える。（隣接する数字が２つある場合は大きい方を評価）
	if triangle[triangle_i-1][item_i-1] < triangle[triangle_i-1][item_i] {
		return triangle[triangle_i-1][item_i]
	} else {
		return triangle[triangle_i-1][item_i-1]
	}
}

// スライス内で一番大きい数を評価
func maxInSlice(slice []int) int {
	max := slice[0]
	for _, num := range slice {
		max = int(math.Max(float64(num), float64(max)))
	}
	return max
}
