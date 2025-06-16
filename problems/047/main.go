package main

import (
	"fmt"
	"slices"
)

func main() {
	solve()
}

// 異なる4つの素因数を持つ最初の4連続の数.
// 4つの素因数を持つかどうかの関数.
func solve() {
	counter := 2
	for {
		if isFourDistinctPrimeNumber(counter) {
			if isFourDistinctPrimeNumber(counter + 1) {
				if isFourDistinctPrimeNumber(counter + 2) {
					if isFourDistinctPrimeNumber(counter + 3) {
						break
					}
				}
			}
		}
		counter++
	}
	fmt.Println(counter)
}

// 素因数の数が4つかを判定
// 本来は素数の場合を考慮する必要があるが、4つの素因数かどうかを確認すればいいので正確な真約数を求めない
func isFourDistinctPrimeNumber(num int) bool {
	x := 2
	numbers := []int{}
	for x <= num {
		if num%x == 0 {
			if !slices.Contains(numbers, x) {
				numbers = append(numbers, x)
			}
			num /= x
			continue
		}
		x++
	}
	return len(numbers) == 4
}
