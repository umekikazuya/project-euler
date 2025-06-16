package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	solve()
}

// Fn =Fn−1 +Fn−2
func solve() {
	x, y := big.NewInt(1), big.NewInt(1)
	count := 2
	for !isPandigital(y.String()) {
		x, y = y, x.Add(x, y)
		count++
		fmt.Println(count)
	}
	fmt.Println(count, y)
}

func isPandigital(target string) bool {
	if len(target) <= 8 {
		return false
	}
	targetForward, targetBehind := target[:9], target[len(target)-9:]
	for i := 1; i <= 9; i++ {
		if !strings.Contains(targetForward, string('0'+byte(i))) || !strings.Contains(targetBehind, string('0'+byte(i))) {
			return false
		}
	}
	return true
}
