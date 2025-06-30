package main

import (
	"fmt"
)

func main() {
	solve()
}

func solve() {
	pentagonalNums := make([]int, 0, 10000)
	cache := make(map[int]struct{}, 10000)
	for i := 1; i <= 10000; i++ {
		pentagonalNums = append(pentagonalNums, getPentagonalNumber(i, cache))
	}
	for x := 0; x < len(pentagonalNums); x++ {
		for y := x + 1; y < len(pentagonalNums); y++ {
			if _, ok := cache[pentagonalNums[x]+pentagonalNums[y]]; ok {
				if _, ok := cache[pentagonalNums[y]-pentagonalNums[x]]; ok {
					fmt.Println(pentagonalNums[y] - pentagonalNums[x])
					return
				}
			}
		}
	}
}

// 五角数を取得する関数
func getPentagonalNumber(n int, cache map[int]struct{}) int {
	x := (3*n*n - n) / 2
	cache[x] = struct{}{}
	return x
}
