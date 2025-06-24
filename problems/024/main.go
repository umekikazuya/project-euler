package main

import (
	"fmt"
)

func main() {
	solve()
}

// 979598, 1217 ns/op
func solve() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 1000000

	result := []int{}

	for 0 < len(arr) {
		x := len(arr)
		all := 1
		// 総パターン数を求める
		for 0 < x {
			all *= x
			x--
		}
		// 先頭を固定した際のパターン数を求める
		dars := all / len(arr)
		y := 0
		index := dars
		for i := 1; i < len(arr); i++ {
			y += dars
			if target <= y {
				break
			}
			index = i
		}
		arr, result = getNumbers(arr, index, result)
		target -= dars * index
	}
	fmt.Println(result)
}

// インデックスから数字を割り当ててその数を引く感じ
func getNumbers(arr []int, index int, result []int) ([]int, []int) {
	if len(arr) <= 1 {
		result = append(result, arr...)
		return arr[:0], result
	}
	result = append(result, arr[index])
	arr = append(arr[:index], arr[index+1:]...)
	return arr, result
}
