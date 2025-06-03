package main

import (
	"fmt"
	"math/big"
)

func main() {
	index := 3
	fibonacciNumbers := []*big.Int{big.NewInt(1), big.NewInt(1)}
	for isThousandDigitNumber(*fibonacciNumbers[len(fibonacciNumbers)-1]) {
		addNumber := addNumberToFibonacci(fibonacciNumbers)
		fibonacciNumbers = append(fibonacciNumbers[:0], fibonacciNumbers[1:]...)
		fibonacciNumbers = append(fibonacciNumbers, addNumber)
		index++
	}
	fmt.Println(index - 1)
}

// フィボナッチ数列に関する処理
// 数列の2個前の数字と1前の数字を足したものを返却
// 内部処理: Fibonacciスライスを受けとって、一番後ろと後ろから二番目の数字を足した数字を返却
func addNumberToFibonacci(arr []*big.Int) *big.Int {
	num := big.NewInt(0)
	return num.Add(arr[len(arr)-1], arr[len(arr)-2])
}

// 1000桁以上の数字かを判定する関数
func isThousandDigitNumber(num big.Int) bool {
	return len(num.String()) <= 999
}
