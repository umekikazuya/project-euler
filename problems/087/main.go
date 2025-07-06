package main

import "fmt"

func main() {
	solve()
}

// 18	  63056567 ns/op
// 18	  65768215 ns/op
func solve() {
	limit := 50000000
	getPrimeNumbers(limit)
}

func getPrimeNumbers(limit int) {
	x := 2
	squereLimit := []int{}
	cubeLimit := []int{}
	fourthPowerLimit := []int{}

	for x*x <= limit {
		if !isPrimeNumber(x) {
			x++
			continue
		}
		if x*x*x*x <= limit {
			fourthPowerLimit = append(fourthPowerLimit, x)
		}
		if x*x*x <= limit {
			cubeLimit = append(cubeLimit, x)
		}
		squereLimit = append(squereLimit, x)
		x++
	}
	result := make(map[int]struct{}, 0)
	for _, fourthPowerLimitNum := range fourthPowerLimit {
		for _, cubeLimitNum := range cubeLimit {
			for _, squereLimitNum := range squereLimit {
				sum := fourthPowerLimitNum*fourthPowerLimitNum*fourthPowerLimitNum*fourthPowerLimitNum + cubeLimitNum*cubeLimitNum*cubeLimitNum + squereLimitNum*squereLimitNum
				if sum < limit {
					result[sum] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(result))
}

func isPrimeNumber(n int) bool {
	x := 2
	for x*x <= n {
		if n%x == 0 {
			return false
		}
		x++
	}
	return true
}
