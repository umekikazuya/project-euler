package main

import (
	"fmt"
)

func main() {
	solve()
}

func solve() {
	const LIMIT = 1000000
	cache := make(map[int]bool)
	// 素数の配列を生成
	primeNumbers := []int{}
	for i := 2; i < LIMIT; i++ {
		if isPrimeNumber(i, cache) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	// 最大長
	// 計算量を減らすために 素数配列の1/3の長さから減らして計測
	k := (len(primeNumbers)+len(primeNumbers)%3)/3 - 1
	for 0 < k {
		windowSize := 0
		for i := 0; i < k; i++ {
			windowSize += primeNumbers[i]
		}
		if LIMIT <= windowSize {
			k--
			continue
		}
		if isPrimeNumber(windowSize, cache) && windowSize < LIMIT {
			fmt.Println(windowSize)
			return
		}

		for i := k; i < len(primeNumbers); i++ {
			windowSize += primeNumbers[i] - primeNumbers[i-k]
			if isPrimeNumber(windowSize, cache) && windowSize < LIMIT {
				fmt.Println(windowSize, k)
				return
			}
		}
		k--
	}
}

func isPrimeNumber(n int, cache map[int]bool) bool {
	if v, ok := cache[n]; ok {
		return v
	}
	x := 2
	for x*x <= n {
		if n%x == 0 {
			cache[n] = false
			return false
		}
		x++
	}
	cache[n] = true
	return true
}
