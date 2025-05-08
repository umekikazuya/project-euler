package main

import (
	"fmt"
)

// フィボナッチ数列を取得
func getFibonacci(below int) []int {
	fibonacci := []int{1, 2}
	for {
		// Belowより末尾の数字が大きい場合はループ終了, 末尾の数字を削除後、配列を返却
		if below < fibonacci[len(fibonacci)-1] {
			return fibonacci[:len(fibonacci)-1]
		}
		// 配列の最後と最後から1つ前を加算したものを配列の末尾にAppend
		fibonacci = append(fibonacci, fibonacci[len(fibonacci)-1]+fibonacci[len(fibonacci)-2])
		continue
	}
}

func main() {
	// 偶数のみ
	sum := 0
	for _, v := range getFibonacci(4000000) {
		if v%2 != 0 {
			continue
		}
		sum += v
	}
	fmt.Println(sum)
}
