package main

import (
	"fmt"
	"math/big"
)

func main() {
	solve()
}

func solve() {
	var sum = 0
	for a := 2; a < 100; a++ {
		if a%10 == 0 {
			continue
		}
		for b := 2; b < 100; b++ {
			result := sumLargeNumber(big.NewInt(1).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil).Text(10))
			if sum < result {
				sum = result
			}
		}
	}
	fmt.Println(sum)
}

// 数字を文字列分解して和を求める関数
func sumLargeNumber(num string) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += int(num[i] - '0')
	}
	return sum
}
