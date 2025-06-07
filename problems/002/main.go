package main

import "fmt"

// だいぶ前だったから忘れたので二回目
func main() {
	fibonacciNumbers := []int{1, 1}
	result := 0
	sum := 0
	for sum < 4000000 {
		sum = fibonacciNumbers[len(fibonacciNumbers)-1] + fibonacciNumbers[len(fibonacciNumbers)-2]
		if sum%2 == 0 {
			result += sum
		}
		fibonacciNumbers = append(fibonacciNumbers, sum)
	}
	fmt.Println(result)
}
