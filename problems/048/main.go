package main

import (
	"fmt"
	"math/big"
)

func main() {
	solve()
}

func solve() {
	result := big.NewInt(0)
	for i := 1; i < 1000; i++ {
		x := big.NewInt(1)
		for j := 1; j <= i; j++ {
			x = x.Mul(x, big.NewInt(int64(i)))
		}
		result = result.Add(result, x)
	}
	fmt.Println(result.Text(10)[len(result.Text(10))-10:])
}
