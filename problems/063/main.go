package main

import (
	"fmt"
	"math/big"
)

func main() {
	solve()
}

func solve() {
	// 桁数, スタートする値
	x, y := int64(2), int64(2)
	// 1乗の「1~9」を初期化入力
	count := 9
	for y < 10 {
		for i := int64(y); i < 10; i++ {
			if len(big.NewInt(1).Exp(big.NewInt(i), big.NewInt(x), nil).Text(10)) == int(x) {
				count++
			} else {
				y = i + 1
			}
		}
		x++
	}
	fmt.Println(count)
}
