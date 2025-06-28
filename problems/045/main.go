package main

import (
	"fmt"
	"math"
)

func main() {
	solve()
}

func solve() {
	var num int64 = 0
	// 五角数285番目の次の条件を確認
	var x int64 = 285
	// 六角数の値を格納
	result := make(map[int64]bool, math.MaxInt8)

	for {
		// 五角数
		if condition(result, (3*x*x-x)/2) {
			num = (3*x*x - x) / 2
			break
		}
		// 六角数
		result[2*x*x-x] = true
		x++
	}
	fmt.Println(num)
}

// 対象の数字が三角数と六角数であることを確認
func condition(result map[int64]bool, target int64) bool {
	if v, ok := result[target]; ok {
		if isTriangularNumber(target) {
			return v
		}
	}
	return false
}

// 三角数判定
func isTriangularNumber(num int64) bool {
	return (num%3 == 0) || num%9 == 1
}
