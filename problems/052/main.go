package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	solve()
}

// 7, 149705649 ns/op
// 19, 57176480 ns/op
// 67, 17395175 ns/op
// 6*xで桁数が変わるやつをスキップしたい.
// 20*6で桁数が変わる.
// 先頭二文字が17以下。
func solve() {
	var limit = 6
	x := 1

	for {
		if !checkCondition(x) {
			x++
			continue
		}
		if checkSameNums(x, limit) {
			break
		}
		x++
	}
	fmt.Println(x)
}

func checkCondition(num int) bool {
	s := strconv.Itoa(num)
	if len(s) < 2 {
		return false
	}
	// 先頭二文字が17以下である場合対象とする(正確には166666666...)。
	target, _ := strconv.Atoi(s[:2])
	return target <= 16
}

// 文字列を辞書順並べ替え
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func checkSameNums(x int, limit int) bool {
	s := sortString(strconv.Itoa(x))
	for i := 2; i <= limit; i++ {
		if s != sortString(strconv.Itoa(x*i)) {
			return false
		}
	}
	return true
}
