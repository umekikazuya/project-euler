package main

import (
	"fmt"
)

const (
	AnswerFormat = "答えは%dです"
	LIMIT        = 100000000
)

func main() {
	fmt.Printf(AnswerFormat, solve())
}

func solve() int {
	primes := generatePrimes()
	return countSpecialCompositeNumber(primes)
}

// generatePrimes は素数のスライスを生成
func generatePrimes() []bool {
	primes := make([]bool, LIMIT)
	primes[0] = true
	primes[1] = true
	for i := 2; i < LIMIT; i++ {
		if primes[i] {
			continue
		}
		for j := i * i; j < LIMIT; j += i {
			primes[j] = true
		}
	}
	return primes
}

func countSpecialCompositeNumber(primes []bool) int {
	count := 0
	for p, isNotPrimeP := range primes {
		if isNotPrimeP {
			continue
		}
		// p, q両方がLIMITの平方根まで行ったらループ終了
		if p > 10000 {
			return count
		}
		count += nestCountSpecialCompositeNumber(primes, p)
	}
	return count
}

func nestCountSpecialCompositeNumber(primes []bool, p int) int {
	count := 0
	for q := p; q < len(primes); q++ {
		if !primes[q] && p*q <= LIMIT {
			count++
		}
		if p*q > LIMIT {
			return count
		}
	}
	return count
}
