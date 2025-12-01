package main

import (
	"slices"
)

const (
	Limit = 1000000
)

func main() {
	solve()
}

func solve() int {
	results := []int{}
	for i := 1; i <= Limit; i++ {
		if _, l, nums := check(i, 0, []int{}); l > len(results) {
			results = append([]int{}, nums...)
		}
	}
	r := Limit
	for _, n := range results {
		if n < r {
			r = n
		}
	}
	return r
}

// check はAmicable Chainsの連鎖の数をカウントする再帰関数
// 連鎖終了時の約数の和と連鎖数を返却
func check(n int, count int, nums []int) (int, int, []int) {
	if isPrime(n) {
		return 0, 0, nil
	}
	if Limit < n {
		return 0, 0, nil
	}
	if len(nums) > 0 && n == nums[0] {
		return n, count, nums
	}
	if slices.Contains(nums, n) {
		return 0, 0, nil
	}
	nums = append(nums, n)
	return check(sumDivisor(n), count+1, nums)
}

// isPrime は素数判定を行う
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	x := 2
	for x*x <= n {
		if n%x == 0 {
			return false
		}
		x++
	}
	return true
}

// sumDivisor は約数の和を算出(0, 1は考慮しない)
func sumDivisor(n int) int {
	sum := 1
	divideNumber := 2
	for divideNumber*divideNumber <= n {
		if n%divideNumber == 0 {
			if divideNumber*divideNumber == n {
				sum += divideNumber
			} else {
				sum += divideNumber
				sum += n / divideNumber
			}
		}
		divideNumber++
	}
	return sum
}
