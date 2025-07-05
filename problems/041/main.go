package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	solve()
}

func solve() {
	x := 987654321
	fmt.Println(isPandigital(x))
	for 0 < x {
		if isPandigital(x) && isPrimeNumber(x) {
			fmt.Println(x)
			return
		}
		x--
	}
}

func isPrimeNumber(n int) bool {
	x := 2
	for x*x <= n {
		if n%x == 0 {
			return false
		}
		x++
	}
	return true
}

func isPandigital(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)
	for i := 1; i <= length; i++ {
		if !strings.Contains(s, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}
