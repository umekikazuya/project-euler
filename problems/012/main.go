package main

import "fmt"

func main() {
	const MAX int64 = 500
	getTriangularNumber(MAX)
}

// 三角数を求め続ける関数
func getTriangularNumber(max int64) {
	var x int64 = 0
	var n int64 = 0

	for countFactor(x) < max {
		n = n + 1
		x = x + n
	}
	fmt.Println(x)
}

// 約数の数を求める関数
func countFactor(num int64) int64 {
	var x int64 = 1
	var count int64 = 0

	// 奇数は捨てる
	if num%2 != 0 {
		return 0
	}
	// 3の倍数は捨てる
	if num%3 != 0 {
		return 0
	}
	// 5の倍数もさすがに捨てていいよね...？
	if num%5 != 0 {
		return 0
	}

	for x <= num {
		if num%x == 0 {
			count = count + 1
		}
		x = x + 1
	}
	return count
}
