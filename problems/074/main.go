package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("答えは%dです", solve())
}

const (
	max      = 1000000
	chainLen = 60
)

func solve() int32 {
	cache := generateFactorialMap()
	count := int32(0)
	s := []int32{60}
	for i := int32(1); i < max; i++ {
		if _, l, _, _ := getChainLength(i, 0, s, cache); l >= chainLen {
			count++
		}
		s = s[:0]
	}
	return count
}

func getChainLength(n int32, count int32, nums []int32, cache map[int32]int32) (int32, int32, []int32, map[int32]int32) {
	if slices.Contains(nums, n) {
		slices.Sort(nums)
		unique := slices.Compact(nums)
		return 0, int32(len(unique)), nil, cache
	}
	// 60以上のチェーンは問題の対象外
	if count >= chainLen+1 {
		return 0, 0, nil, cache
	}
	nums = append(nums, n)
	return getChainLength(getFactorialSum(n, cache), count+1, nums, cache)
}

func generateFactorialMap() map[int32]int32 {
	cache := make(map[int32]int32, 9)
	// 数学的な整合性(ルール)
	cache[0] = int32(1)
	cache[1] = int32(1)
	sum := int32(1)
	for i := int32(2); i <= 9; i++ {
		sum *= i
		cache[i] = sum
	}
	return cache
}

func getFactorialSum(n int32, cache map[int32]int32) int32 {
	div := n
	sum := int32(0)
	for {
		if w, ok := cache[int32(div%10)]; ok {
			sum += w
		}
		if div < 10 {
			break
		}
		div = div / 10
	}
	return sum
}
