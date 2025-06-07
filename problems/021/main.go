package main

import "fmt"

func main() {
	sum := 0
	cache := make(map[int]int)
	for i := 2; i < 10000; i++ {
		if isFraternalNumber(i, cache) {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}

// 友愛数かどうかを求める関数
func isFraternalNumber(a int, cache map[int]int) bool {
	b := getFactorSum(a, cache)
	if a == b {
		return false
	}
	c := getFactorSum(b, cache)
	return a == c
}

// 約数の和を求める関数
func getFactorSum(num int, cache map[int]int) int {
	if v, ok := cache[num]; ok {
		return v
	}
	sum := 1
	x := 2
	for x*x <= num {
		if num%x == 0 {
			sum = sum + x
			if x*x != num {
				sum = sum + num/x
			}
		}
		x++
	}
	cache[num] = sum
	return sum
}
