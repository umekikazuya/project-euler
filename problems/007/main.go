package main

import "fmt"

func main() {
	count := 1
	var num int32 = 2
	for count <= 10001 {
		if isPrimeNumber(num) {
			count++
		}
		num++
	}
	fmt.Println(num - 1)
}

func isPrimeNumber(num int32) bool {
	var x int32 = 2
	for x*x <= num {
		if num%x == 0 {
			return false
		}
		x++
	}
	return true
}
