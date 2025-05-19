package main

import "fmt"

func main() {
	var max int = 2000000

	sum := 0
	for i := 2; i < max; i++ {
		if isPrimeNumber(i) {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}

func isPrimeNumber(x int) bool {
	var y int = 2
	for y*y <= x {
		if x%y == 0 {
			return false
		}
		y = y + 1
	}
	return true
}
