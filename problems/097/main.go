package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Printf("答えは...。%sです。", solve())
}

func solve() string {
	n := 10
	s := calculate()
	result := s[len(s)-n:]
	return result
}

func calculate() string {
	a, _ := new(big.Int).SetString("28433", 10)
	return new(big.Int).Add(new(big.Int).Mul(a, new(big.Int).Exp(big.NewInt(2), big.NewInt(7830457), nil)), big.NewInt(1)).String()
}
