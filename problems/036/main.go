package main

import (
	"fmt"
	"strconv"
)

func main() {
	solve()
}

func solve() {
	sum := 0
	for i := 1; i < 1000000; i++ {
		if isFrequency(strconv.Itoa(i)) && isFrequency(strconv.FormatInt(int64(i), 2)) {
			sum += i
		}
	}
	fmt.Println(sum)
}

// 回文数判定
func isFrequency(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) == s
}
