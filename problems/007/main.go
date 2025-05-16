package main

import "fmt"

func main() {
	const TARGET = 10001

	var count int32 = 0
	var current int32 = 2
	for count < TARGET {
		if isPrimeNumber(current) {
			count = count + 1
		}
		current = current + 1
	}
	fmt.Println(current - 1)
}

// 素数チェック
func isPrimeNumber(x int32) bool {
	var y int32 = 2
	for y < x {
		if x%y == 0 {
			return false
		}
		y = y + 1
	}
	return true
}
