package main

import (
	"fmt"
	"strconv"
)

func main() {
	solve()
}

func solve() {
	count := 0
	for i := 11; i <= 10000; i++ {
		if isLychrelNumber(i) {
			count++
		}
	}
	fmt.Println(count)
}

// 受け取った数字を前後入れ替えて和を求める
// その和が回分数かを確認する
func isLychrelNumber(num int) bool {
	for i := 1; i <= 50; i++ {
		reverseNum, _ := strconv.Atoi(reverseString(strconv.Itoa(num)))
		// 元の数字と入れ替えた数字を足す
		num += reverseNum
		// 和が回分数かチェックする
		reverseSum, _ := strconv.Atoi(reverseString(strconv.Itoa(num)))
		if num == reverseSum {
			return false
		}
	}
	return true
}

// 順番を入れ替える
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
