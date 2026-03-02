package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solve())
}

const LIMIT = 10000000

func solve() int {
	arr := generateDivisorsNumbers()
	count := 0
	for i := 2; i <= LIMIT; i++ {
		if arr[i] == arr[i+1] {
			count++
		}
	}
	return count
}

func generateDivisorsNumbers() []int {
	arr := make([]int, LIMIT+2)
	for p := 1; p <= LIMIT+1; p++ {
		for i := p; i <= LIMIT+1; i += p {
			arr[i]++
		}
	}
	return arr
}

func numuberOfDivisors(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			count += 2
		}
	}
	if sqrtN*sqrtN == n {
		count--
	}
	return count
}
