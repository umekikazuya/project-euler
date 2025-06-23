package main

import (
	"fmt"
	"math/big"
)

func main() {
	solve()
}

func solve() {
	x, y := big.NewInt(1), big.NewInt(1)
	count := 2
	for len(y.Text(10)) < 1000 {
		x, y = y, x.Add(x, y)
		count++
	}
	fmt.Println(count)
}
