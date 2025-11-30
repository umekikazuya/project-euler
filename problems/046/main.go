package main

import (
	"math"
)

func main() {
	solve()
}

// 合成数を算出する
// 素数と平方数の2倍の和=奇数の合成数、x + 2y^2 = z
// 対象: 全ての奇数の合成数
func solve() int {
	var max int = 10000
	numbers := getNumbers(max)
	for num, isPrime := range numbers {
		// 1, 2は除外
		if num <= 2 {
			continue
		}
		if num%2 == 0 {
			continue
		}
		// 素数は除外
		if isPrime {
			continue
		}
		if !checkGoldbach(num, numbers) {
			if num < max {
				max = num
			}
		}
	}
	return max
}

// getNumbers
func getNumbers(max int) map[int]bool {
	nums := make(map[int]bool, max)
	for i := int(0); i <= max; i++ {
		nums[i] = true
	}
	nums[0] = false
	nums[1] = false
	for i := int(2); i <= max; i++ {
		for j := i * i; j <= max; j += i {
			nums[j] = false
		}
	}
	return nums
}

func isPrime(n int, cache map[int]bool) bool {
	if v, ok := cache[n]; ok {
		return v
	}
	var a int = 2
	if n < a {
		return false
	}
	for a*a <= n {
		if n%a == 0 {
			return false
		}
		a++
	}
	return true
}

// checkGoldbach はChristian Goldbachの提唱通りかを検証
// 素数と平方数の2倍の和=奇数の合成数、primeNumber + 2y^2 = compositeNumber
func checkGoldbach(compositeNumber int, cache map[int]bool) bool {
	for i := 1; i <= int(getSqrt(compositeNumber)); i++ {
		if isPrime(
			int(compositeNumber-2*int(math.Pow(float64(i), 2))),
			cache,
		) {
			return true
		}
	}
	return false
}

// getSqrt は平方根を求める
func getSqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}
