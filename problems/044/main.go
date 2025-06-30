package main

import (
	"fmt"
	"slices"
)

func main() {
	solve()
}

func solve() {
	pentagonalNums := []int{}
	for i := 1; i <= 10000; i++ {
		pentagonalNums = append(pentagonalNums, getPentagonalNumber(i))
	}
	for x := 0; x < len(pentagonalNums); x++ {
		for y := x + 1; y < len(pentagonalNums); y++ {
			if slices.Contains(pentagonalNums, pentagonalNums[x]+pentagonalNums[y]) && slices.Contains(pentagonalNums, pentagonalNums[y]-pentagonalNums[x]) {
				fmt.Println(pentagonalNums[y] - pentagonalNums[x])
				return
			}
		}
	}
}

// 五角数を取得する関数
func getPentagonalNumber(n int) int {
	return (3*n*n - n) / 2
}
