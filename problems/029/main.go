package main

import (
	"fmt"
	"math/big"
)

func main() {
	solve()
}

// 2~100
func solve() {
	a, b := 2, 100
	num := make(map[string]bool)
	for i := a; i <= b; i++ {
		x := big.NewInt(int64(i))
		for j := a; j <= b; j++ {
			x := x.Mul(x, big.NewInt(int64(i)))
			if !num[x.Text(10)] {
				num[x.Text(10)] = true
			}
		}
	}
	fmt.Println(len(num))
}
