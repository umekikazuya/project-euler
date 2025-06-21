package main

import (
	"fmt"
)

func main() {
	solve()
}

func solve() {
	count := 2
	sum := 1
	// 1ループ増えるごとに+2ずつ加算される
	roopCount := 0
	for count <= 1001*1001 {
		for i := 0; i < 4; i++ {
			count += getAdvanceNum(i, roopCount)
			sum += count
			// 4つ目の角の場合は-1をする
			if i == 3 {
				sum--
			}
		}
		roopCount = roopCount + 2
	}
	fmt.Println(sum)
}

// スパイラルの向きに応じて進める数を取得
func getAdvanceNum(num int, roopCount int) int {
	mapping := map[int]int{
		0: 1,
		1: 2,
		2: 2,
		3: 3,
	}
	return mapping[num] + roopCount
}
