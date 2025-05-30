package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 145は興味深い数字です。なぜなら、
// 1! + 4! + 5! = 1 + 24 + 120 = 145 となり、自身の各桁の階乗の和と等しくなるからです。
// 各桁の階乗の和と等しくなるすべての数の合計を求めてください。
// ※注意：1! = 1 および 2! = 2 は「和」ではないため、含めません。
func main() {
	baseNum := 3
	sum := 0
	cache := make(map[int]int)

	// 期待を込めて1,000,000で処理を止める
	// 恐らく9!が362,880なので、どうがんばっても1,000,000には届かない
	for baseNum < 1000000 {
		splitNumbers := splitNumberToSlice(baseNum)
		if baseNum == sumFactorialNumbers(splitNumbers, cache) {
			sum = sum + baseNum
		}
		baseNum++
	}
	fmt.Println(sum)
}

// 数字を一文字ずつに分割してスライスに変換する関数
func splitNumberToSlice(num int) []int {
	a := strconv.Itoa(num)
	splitStrings := strings.Split(a, "")
	var splitNumbers []int
	for _, b := range splitStrings {
		c, _ := strconv.Atoi(b)
		splitNumbers = append(splitNumbers, c)
	}
	return splitNumbers
}

// スライス内の数字の階乗の合計を返す関数
func sumFactorialNumbers(arr []int, cache map[int]int) int {
	sum := 0
	for _, num := range arr {
		sum = sum + getFactorialNumber(num, cache)
	}
	return sum
}

// 階乗を求める関数
func getFactorialNumber(num int, cache map[int]int) int {
	result := 1
	if v, ok := cache[num]; ok {
		return v
	}
	for i := 1; i <= num; i++ {
		result = result * i
	}
	cache[num] = result
	return result
}
