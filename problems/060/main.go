package main

import (
	"fmt"
	"strconv"
)

func main() {
	solve()
}

func solve() {
	primeNumbers := getPrimeNumbers(10000)
	index := 0
	goal := [][]int{}
	for index < len(primeNumbers) {
		current := []int{primeNumbers[index]}

		i := index + 1
		goal = updateNumbersSlice(current, primeNumbers, i, &goal)
		index++
	}

	target := 1000000
	for _, nums := range goal {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		if sum < target {
			target = sum
		}
	}
	fmt.Println(target)
}

// isPrimeNumber は素数判定を行う
func isPrimeNumber(x int) bool {
	a := 2
	for a*a <= x {
		if x%a == 0 {
			return false
		}
		a++
	}
	return true
}

// getPrimeNumbers は上限を引数で受け取ってその数字までの素数をスライスで取得する
func getPrimeNumbers(max int) []int {
	x := int(2)
	primeNumbers := []int{}
	for x < max {
		if isPrimeNumber(int(x)) {
			primeNumbers = append(primeNumbers, int(x))
		}
		x++
	}
	return primeNumbers
}

// 2つの数字を連結する
func concatenateNumbers(a, b int) (int, int) {
	result1 := a*(digit(b)) + b
	result2 := b*(digit(a)) + a
	return result1, result2
}

func digit(x int) int {
	a := 1
	for i := 0; i < len(strconv.Itoa(x)); i++ {
		a *= 10
	}
	return a
}

// スライスを受け取って、素数軍団(5個)を形成する再帰関数
func updateNumbersSlice(current []int, list []int, index int, results *[][]int) [][]int {
	if len(current) == 5 {
		*results = append(*results, current)
		return *results
	}
	for i := index; i < len(list)-1; i++ {
		if !check(current, list[i+1]) {
			continue
		}
		copyCurrent := append([]int(nil), current...)
		copyCurrent = append(copyCurrent, list[i+1])
		updateNumbersSlice(copyCurrent, list, i+1, results)
	}
	return *results
}

// check は第一引数のスライス内の素数と第二引数の素数を検証する
func check(current []int, target int) bool {
	for _, num := range current {
		got, got1 := concatenateNumbers(num, target)
		if !isPrimeNumber(got) {
			return false
		}
		if !isPrimeNumber(got1) {
			return false
		}
	}
	return true
}
