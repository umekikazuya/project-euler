package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 循環素数
// 1. 1000000までを回す
// 2. 循環素数であれば配列に追加
func main() {
	// 上限
	const LIMIT int = 1000000
	// 素数配列
	primeNumbers := []int{}
	// 上限までループ処理
	for i := 2; i < LIMIT; i++ {
		if isCircularPrimeNumber(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	fmt.Println(len(primeNumbers))
}

// 素数判定(割り切れる場合をTRUE)
func isPrimeNumber(num int) bool {
	x := 2
	for x*x <= num {
		if num%x == 0 {
			return false
		}
		x++
	}
	return true
}

// 循環素数判定(引数は素数のみを受け付ける想定、割り切れる場合をTRUEで返却)
func isCircularPrimeNumber(num int) bool {
	arr := strings.Split(strconv.Itoa(num), "")
	for i := range arr {
		rotated := append(arr[i:], arr[:i]...)
		rotatedNumber, _ := strconv.Atoi(strings.Join(rotated, ""))
		if !isPrimeNumber(rotatedNumber) {
			return false
		}
	}
	return true
}
