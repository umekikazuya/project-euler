package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cache := make(map[int]int)
	max := getSumOfFifthPower(9, cache)
	length := len(strconv.Itoa(max)) + 1
	result := 0

	// 1は加算対象外のため2スタート
	for i := 2; i <= max*length; i++ {
		if i == getSumOfFifthPower(i, cache) {
			result += i
		}
	}
	fmt.Println(result)
}

// 5桁の数字を1文字ずつの文字列に分解し、分解した数の5乗を加算する
func getSumOfFifthPower(num int, cache map[int]int) int {
	x := strconv.Itoa(num)
	arr := strings.Split(x, "")

	result := 0
	for _, item := range arr {
		y, _ := strconv.Atoi(item)
		result += getFifthPower(y, cache)
	}
	return result
}

// 5乗を求める関数
func getFifthPower(num int, cache map[int]int) int {
	if v, ok := cache[num]; ok {
		return v
	}
	x := num * num * num * num * num
	cache[num] = x
	return x
}
