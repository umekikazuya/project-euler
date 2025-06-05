package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	target := getFactorial(100)

	result := digitSum(target)
	fmt.Println(result)
}

// 数字を一文字ずつに分割して加算する関数
func digitSum(number *big.Int) int {
	arr := strings.Split(number.Text(10), "")
	sum := 0
	for _, item := range arr {
		num, _ := strconv.Atoi(item)
		sum = sum + num
	}
	return sum
}

// 階乗を求める関数
func getFactorial(num int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= num; i++ {
		result = new(big.Int).Mul(result, big.NewInt(int64(i)))
	}
	return result
}
