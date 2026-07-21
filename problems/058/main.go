package main

import (
	"fmt"
)

func main() {
	fmt.Printf("答え: %d", solve())
}

func solve() int {
	per := 100
	totalCount := 1
	primeNumCount := 0 // 対角線を辿っていく中での素数の総数
	current := 1
	loop := 1
	addNumber := 2
	for per >= 10 {
		for i := 0; i < 4; i++ {
			current = current + addNumber
			if isPrime(current) {
				primeNumCount++
			}
			per = primeNumCount * 100 / totalCount
			totalCount++
		}
		addNumber = addNumber + 2
		loop++
	}
	return addNumber - 3
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	x := 2
	for x*x <= num {
		if num%x == 0 {
			return false
		}
		x++
	}
	return true
}
