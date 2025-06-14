package main

import (
	"fmt"
	"strconv"
)

// 11個のTruncatable Primesを求める
func main() {
	num := 11
	result := 0
	count := 0
	for count < 11 {
		if isPrimeNumber(num) && isTruncatablePrimeNumber(num) {
			result += num
			count++
		}
		num++
	}
	fmt.Println(result)
}

func isPrimeNumber(num int) bool {
	if num == 1 {
		return false
	}
	x := 2
	for x*x <= num {
		if num%x == 0 {
			return false
		}
		x++
	}
	return true
}

// 1文字削除していって、切り取られた数字が素数(Truncatable Primes)か判定する関数
// 素数じゃない時点でアーリーリターン
func isTruncatablePrimeNumber(num int) bool {
	target := (strconv.Itoa(num))
	for i := 0; i < len(target); i++ {
		// 左から
		targetNum, _ := strconv.Atoi(target[(i + 1):])
		// 右から
		targetReverseNum, _ := strconv.Atoi(target[:len(target)-1-i])
		if !isPrimeNumber(targetNum) {
			return false
		}
		if !isPrimeNumber(targetReverseNum) {
			return false
		}
	}

	return true
}
