package main

import "fmt"

// 13195の素因数は5、7、13、29である。
// 600851475143の最大の素因数は何か。
func main() {
	result := []int{}
	for _, num := range getFactor(600851475143) {
		if isPrimeNumber(num) {
			result = append(result, num)
		}
	}
	fmt.Println(result)
}

// 約数の取得
func getFactor(num int) []int {
	result := []int{}
	x := 2
	for x*x <= num {
		if num%x == 0 {
			result = append(result, x)
			if x*x != num {
				result = append(result, num/x)
			}
		}
		x++
	}
	return result
}

// 素数判定
func isPrimeNumber(num int) bool {
	x := 2
	for x*x <= num {
		if num%x == 0 {
			return false
		}
		x++
	}
	return true
}
