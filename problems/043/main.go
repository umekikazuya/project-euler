package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	solve()
}

func solve() {
	getNumbers()
	// for i := int(123456789); i <= 9876543210; i++ {
	// 	if isPandigital(i) {
	// 		fmt.Println(i)
	// 	}
	// }
}

// func isPandigital(num int) bool {
// 	targets := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
// 	arr := strings.Split(strconv.Itoa(int(num)), "")
// 	for _, target := range targets {
// 		if !slices.Contains(arr, target) {
// 			return false
// 		}
// 	}
// 	return true
// }

// 重複がない3桁の条件を満たす数を取得する
// その統合を行い、重複があるないを求めて、パターン数を求める
func getNumbers() {
	primeNumbers := [7]int{2, 3, 5, 7, 11, 13, 17}
	result := make(map[int][]string)
	for _, num := range primeNumbers {
		for i := num; i < 1000; i += num {
			if clearCondition(i) {
				s := strconv.Itoa(i)
				if len(s) < 3 {
					s = "0" + s
				}
				if clearCondition2(s, num) {
					result[num] = append(result[num], s)
				}
			}
		}
	}
	fmt.Println(result)
	for _, nums := range result {
		for _, num := range nums {
			fmt.Println(num)
		}
	}
}

// 3桁の数が問題の条件を満たすかの判定
func clearCondition(n int) bool {
	// 二桁未満は除外(001, 002の場合、0が重複するため)。同様に010, 011は除外。
	if n < 12 {
		return false
	}
	// 二桁で10で割り切れるものは除外(010, 020)
	if n < 100 && n%10 == 0 {
		return false
	}
	return isPandigital(n)
}

// 上二桁が前の素数で割り切れるか判定
func clearCondition2(s string, beforePrimeNumber int) bool {
	if beforePrimeNumber == 2 {
		return true
	}
	arr := strings.Split(s, "")
	newS := arr[0] + arr[1]
	n, _ := strconv.Atoi(newS)
	return n%beforePrimeNumber == 0
}

func isPandigital(n int) bool {
	arr := strings.Split(strconv.Itoa(n), "")
	sort.Strings(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
	}
	return true
}
