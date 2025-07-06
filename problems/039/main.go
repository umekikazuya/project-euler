package main

import "fmt"

func main() {
	solve()
}

func solve() {
	limit := int32(1000)
	counts := make(map[int32]int32, 1000)
	a := int32(1)
	for a <= limit/3 {
		for b := a; b <= limit; b++ {
			p := getSqrt(a*a+b*b) + a + b
			if p <= limit {
				counts[p]++
			}
		}
		a++
	}
	result := int32(0)
	for n, count := range counts {
		if result < count {
			fmt.Println(n)
			result = count
		}
	}
}

// 平方根が整数かを判定
func getSqrt(n int32) int32 {
	x := int32(2)
	for x*x <= n {
		if x*x == n {
			return x
		}
		x++
	}
	return 1000
}
