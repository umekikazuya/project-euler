package main

import (
	"fmt"
)

func main() {
	r := solve()
	fmt.Printf("答えは %v です", r)
}

func solve() int {
	count := 0
	for i := 1; i <= 10000000; i++ {
		if isEightyNineSquareDigit(i) {
			count++
		}
	}
	return count
}

// 任意の数を受け取って、各桁の二乗の和を計算し、89に判定する関数
func isEightyNineSquareDigit(n int) bool {
	a := sumSquareDigit(n)
	if a == 1 {
		return false
	}
	if a == 89 {
		return true
	}
	return isEightyNineSquareDigit(a)
}

// sumSquareDigit は各桁の二乗の和を計算する関数
func sumSquareDigit(n int) int {
	sum := 0
	for n > 0 {
		a := n % 10
		sum += a * a
		n = n / 10
	}
	return sum
}
