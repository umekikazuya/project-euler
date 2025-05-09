package main

import (
	"fmt"
)

func main() {
	primeNumbers := execPrimeFactor(600851475143)
	fmt.Println(primeNumbers)
}

// 素因数分解
func execPrimeFactor(target int) int {
	x := 2
	for x < target {
		if target%x == 0 {
			if isPrimeNumber(target / x) {
				return target / x
			}
		}
		x = x + 1
	}
	return 0
}

// 素数チェック
func isPrimeNumber(x int) bool {
	a := 2
	for a < x {
		if x%a == 0 {
			return false
		}
		a = a + 1
	}
	return true
}
