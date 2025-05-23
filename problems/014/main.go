package main

import "fmt"

// 1. コラッツ関数
// 2.任意の数の連鎖数を求める関数を定義
// 3. 1000000からループ
func main() {
	const TARGET int = 1000000

	// 答え
	var x int = 0
	// 最大の連鎖数
	var y int = 0
	for v := range TARGET {
		var count int = getChainCount(v)
		if y < count {
			x = v
			y = count
		}
	}
	fmt.Println(x)
}

// 任意の数の連鎖数を求める関数
func getChainCount(x int) int {
	var count int = 0
	for 1 < x {
		x = collatz(x)
		count++
	}
	return count
}

// コラッツ関数
func collatz(n int) int {
	// nが偶数の場合: n/2
	if n%2 == 0 {
		return n / 2
	}
	// nが奇数の場合: 3n+1
	return 3*n + 1
}
